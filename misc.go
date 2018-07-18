package sysgosys

func (s *sContext) Whoami() string {
	return s.uname
}
