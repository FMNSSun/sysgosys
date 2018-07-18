package sysgosys

import (
	utils "github.com/FMNSSun/sysgosys/utils.fs"
)

func (s *sContext) Open(fpath string, flgs uint16) (Handle, SError) {
	parts := utils.SplitRoot(fpath)

	if parts == nil {
		return NO_HANDLE, SErrorf(ERR_FS_INVALID_PATH, "Path %q is not a valid path.")
	}

	root := parts[0]
	path := utils.SplitPath(parts[1])

	if s.fss == nil {
		return NO_HANDLE, SErrorf(ERR_FS_NO_SUCH_FS, "No filesystems for this context.")
	}

	fs := s.fss[root]

	if fs == nil {
		return NO_HANDLE, SErrorf(ERR_FS_NO_SUCH_FS, "No such filesystem for this context: %q", root)
	}

	// Check permissions. 
	for i := 0; i < len(path) - 1; i++ {
		di, err := fs.DirInfo(path[:i+1])

		if err != nil {
			return NO_HANDLE, err
		}

		// Does current user have DIR_ENTER permissions?
		dirperms := di.Permissions()[s.uname]

		if (dirperms & PERM_DIR_ENTER) != PERM_DIR_ENTER {
			return NO_HANDLE, SErrorf(ERR_FS_ACCESS_DENIED, "User %q can not enter %q.", fpath, s.uname)
		}
	}

	fi, err := fs.FileInfo(path)

	if err != nil {
		return NO_HANDLE, err
	}

	if fi == nil {
		return NO_HANDLE, SErrorf(ERR_FS_FILE_NOT_FOUND, "File %q does not exist.", fpath)
	}

	fperms := fi.Permissions()

	perms := fperms[s.uname]

	omode := flgs & 0x03

	switch omode {
	case OF_READ_ONLY:
		if (perms & PERM_FILE_READ) != PERM_FILE_READ {
			return NO_HANDLE, SErrorf(ERR_FS_ACCESS_DENIED, "Read access denied to %q for user %q.", fpath, s.uname)
		}
	case OF_WRITE_ONLY:
		if (perms & PERM_FILE_WRITE) != PERM_FILE_WRITE {
			return NO_HANDLE, SErrorf(ERR_FS_ACCESS_DENIED, "Write access denied to %q for user %q.", fpath, s.uname)
		}
	case OF_READ_WRITE:
		if (perms & PERM_FILE_READ) != PERM_FILE_READ {
			return NO_HANDLE, SErrorf(ERR_FS_ACCESS_DENIED, "Read/Write access denied to %q for user %q.", fpath, s.uname)
		}
		if (perms & PERM_FILE_WRITE) != PERM_FILE_WRITE {
			return NO_HANDLE, SErrorf(ERR_FS_ACCESS_DENIED, "Read/write access denied to %q for user %q.", fpath, s.uname)
		}
	}

	// try to alloc a handle
	handle := s.AllocHandle()

	if handle == NO_HANDLE {
		return handle, SErrorf(ERR_FS_ALLOC_HANDLE, "Could not allocate a handle.")
	}

	f, err := fs.OpenFile(path, flgs)

	if err != nil {
		s.ReleaseHandle(handle)
		return NO_HANDLE, err
	}

	if f == nil {
		s.ReleaseHandle(handle)
		return NO_HANDLE, SErrorf(ERR_FS_FILE_NOT_FOUND, "File %q does not exist.", fpath)
	}

	err = s.AssocHandle(handle, f)

	if err != nil {
		f.Abort()
		return NO_HANDLE, err
	}

	return handle, nil
}
