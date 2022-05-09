package handlers

import (
	"context"

	addservice "github.com/evgeniy-dammer/simplevaluesgrpc/src/addservice/proto"
)

type AddServiceServer struct {
	addservice.UnimplementedAddServiceServer
}

func (*AddServiceServer) Add(ctx context.Context, in *addservice.AddRequest) (*addservice.AddResponse, error) {
	return &addservice.AddResponse{Result: in.A + in.B}, nil
}
