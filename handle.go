package sysgosys

func (s *sContext) AllocHandle() Handle {
	s.hndlMutex.Lock()
	defer s.hndlMutex.Unlock()

	it := s.hndl
	s.hndl++

	return Handle(it)
}

func (s *sContext) ReleaseHandle(handle Handle) {
	// TODO
}

func (s *sContext) AssocHandle(handle Handle, with interface{}) {
	s.hndlMutex.Lock()
	defer s.hndlMutex.Unlock()

	_, ok := s.handles[handle]

	if ok {
		// TODO: this should probably panic. 
	}

	s.handles[handle] = with
}
