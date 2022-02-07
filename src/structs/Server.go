package structs

type Server struct {
	URL        string
	ServerName string
}

func (s *Server) NewServer(url, serverName string) {
	s.ServerName = serverName
	s.URL = url
}
