package server

func (s *server) routes() {
	s.router.HandleFunc("/", s.Authenticate(s.rootHandler()))
}
