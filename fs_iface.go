package sysgosys

type Permissions map[string]uint8

const PERM_FILE_READ = uint8(0x01)
const PERM_FILE_WRITE = uint8(0x02)

const PERM_DIR_ENTER = uint8(0x01)
const PERM_DIR_LIST = uint8(0x01)

type FileSys interface {
	OpenFile(fpath []string, flgs uint16) (File, SError)
	OpenDir(fpath []string) (Dir, SError)
	FileInfo(fpath []string) (FileInfo, SError)
	DirInfo(fpath []string) (DirInfo, SError)
}

type File interface {
	Read([]byte) (int, SError)
	Write([]byte) (int, SError)
	Seek(uint64) SError
	Abort()
}

type Dir interface {
	ListFiles() (FileIterator, SError)
	ListDirs() (DirIterator, SError)
}

type DirIterator interface {
	Next() (DirInfo, SError)
}

type FileIterator interface {
	Next() (FileInfo, SError)
}


type DirInfo interface {
	Name() string
	Permissions() Permissions
}

type FileInfo interface {
	Name() string
	Size() uint64
	Permissions() Permissions
}
