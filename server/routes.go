package server

func (s *server) routes() {
	s.router.HandleFunc("/locate", s.loggingMiddleware(s.handleLocate())).Methods("POST")
}
