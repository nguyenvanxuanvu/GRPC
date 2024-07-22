package grpc_handler

import (
	"context"

	"learning/grpc/pkg/model"

	"learning/grpc/util/protoc/pb"
)

type grpcHandler struct {
	userSvc    userService
	pb.UnimplementedUserServer
}

type userService interface {
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
}

func NewGRPCHandler(userSvc userService) *grpcHandler {
	return &grpcHandler{
		userSvc:    userSvc,
	}
}