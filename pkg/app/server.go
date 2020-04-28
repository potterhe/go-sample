package app

import (
	"time"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct {
}

func (s *Server) Serve() {
	time.Sleep(5 * time.Minute)
}
