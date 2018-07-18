package sysgosys

import "sync"

type sContext struct {
	uname string
	fss map[string]FileSys
	hndl uint16
	hndlMutex *sync.Mutex
	handles map[Handle]interface{}
}

