package sysgosys

import "fmt"

type Handle uint16
type ErrorCode uint32

const NO_HANDLE = Handle(0xFFFF)

type SError interface {
	ErrCode() ErrorCode
	Msg() string
}

type sError struct {
	errCode ErrorCode
	msg string
}

func (s *sError) ErrCode() ErrorCode {
	return s.errCode
}

func (s *sError) Msg() string {
	return s.msg
}

func SErrorf(errCode ErrorCode, msg string, args... interface{}) SError {
	return &sError {
		errCode: errCode,
		msg: fmt.Sprintf(msg, args...),
	}
}

type S interface {
	// Read reads up to max(n, len(dstbuf) bytes into dstbuf from the handle.
	// On EOF this returns -1 bytes read. If currently no data is available
	// this returns -2. 
	Read(handle Handle, dstbuf []byte, n int) (int, ErrorCode)

	// Write write up to len(srcbuf) bytes to fd. Returns the amount of
	// bytes written. 
	Write(handle Handle, srcbuf []byte) (int, ErrorCode)

	// Open opens a "file" and returns a Handle. 
	Open(fpath string, flgs uint16) (Handle, ErrorCode)

	// Close closes the handle. 
	Close(handle Handle) ErrorCode

	// Whoami returns the name of the calling user.
	Whoami() string
}
