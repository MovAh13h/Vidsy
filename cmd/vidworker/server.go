package main

import (
	"context"
	"errors"

	"github.com/MovAh13h/Vidsy/proto"
)

type Server struct {
	proto.UnimplementedVidsyServer
}

func (s *Server) Edit(ctx context.Context, req *proto.Request) (*proto.Response, error) {

	return nil, errors.New("Apple")
}