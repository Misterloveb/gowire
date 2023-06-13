package common

import "gindemo/pkg/log"

type Server struct {
	*log.Logger
}

func NewServer(log *log.Logger) *Server {
	return &Server{
		Logger: log,
	}
}
