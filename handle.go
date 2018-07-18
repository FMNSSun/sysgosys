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

func (s *sContext) AssocHandle(handle Handle, with interface{}) SError {
	s.hndlMutex.Lock()
	defer s.hndlMutex.Unlock()

	_, ok := s.handles[handle]

	if ok {
		return SErrorf(ERR_CORE_HANDLE_EXISTS, "Handle %d already exists.", handle)
	}

	s.handles[handle] = with

	return nil
}
