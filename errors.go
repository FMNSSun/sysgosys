package sysgosys

const ERR_SUB_CORE = uint32(uint32(0x01) << 22)
const ERR_SUB_FS = uint32(uint32(0x02) << 22)

const ERR_FS_NOT_SEEKABLE = ErrorCode(ERR_SUB_FS | 0x01)
const ERR_FS_FILE_NOT_FOUND = ErrorCode(ERR_SUB_FS | 0x02)
const ERR_FS_DIR_NOT_FOUND = ErrorCode(ERR_SUB_FS | 0x03)
const ERR_FS_INVALID_PATH = ErrorCode(ERR_SUB_FS | 0x04)
const ERR_FS_NO_SUCH_FS = ErrorCode(ERR_SUB_FS | 0x05)
const ERR_FS_ACCESS_DENIED = ErrorCode(ERR_SUB_FS | 0x06)
const ERR_FS_ALLOC_HANDLE = ErrorCode(ERR_SUB_FS | 0x07)

const ERR_CORE_HANDLE_EXISTS = ErrorCode(ERR_SUB_CORE | 0x01)
