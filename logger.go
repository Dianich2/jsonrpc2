package jsonrpc2

type Logger interface {
	Printf(format string, v ...interface{})
}

func (s *Server) SetLogger(logger Logger) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.logger = logger
}
