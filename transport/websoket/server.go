package websoket

import (
	"context"
	"github.com/olahol/melody"
)

type Config melody.Config

type Server struct {
	*melody.Melody
}

func NewServer(config *Config) *Server {
	return &Server{
		Melody: melody.New(),
	}
}

func (s *Server) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
