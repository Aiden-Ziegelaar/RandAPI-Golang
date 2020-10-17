package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"randAPI/generators"
	"randAPI/proto/grpcService"
)

type server struct {
	grpcService.RandAPIServer

	serverName string
}

func (s *server) Boolean (ctx context.Context, req *empty.Empty) (*grpcService.BooleanResponse, error) {
	var resp = generators.GenerateBoolean()
	return &resp, nil
}

func (s *server) WeightedBoolean (ctx context.Context, req *grpcService.WeightedBooleanRequest) (*grpcService.BooleanResponse, error) {
	var resp = generators.GenerateWeightedBoolean(int(req.FalseWeight), int(req.TrueWeight))
	return &resp, nil
}

func (s *server) Integer (ctx context.Context, req *empty.Empty) (*grpcService.IntegerResponse, error) {
	var resp = generators.GenerateInteger()
	return &resp, nil
}

func (s *server) IntegerN (ctx context.Context, req *grpcService.IntegerRequest) (*grpcService.IntegerResponse, error) {
	var resp = generators.GenerateIntegerN(req.Max)
	return &resp, nil
}

