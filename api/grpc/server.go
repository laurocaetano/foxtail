package grpcserver

import (
	"context"

	pb "github.com/laurocaetano/foxtail/gen/protos/v1/foxtailservice"
)

func Put(ctx context.Context, request *pb.PutRequest) (*pb.PutResponse, error) {
	return &pb.PutResponse{Status: pb.ResponseStatus_RESPONSE_STATUS_OK}, nil
}
