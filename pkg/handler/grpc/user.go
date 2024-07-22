package grpc_handler

import (
	"context"

	"learning/grpc/util/protoc/pb"
)

func (h *grpcHandler) GetUserByID(ctx context.Context, args *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	var (
		err error
	)

	id := args.GetId()
	if id <= 0 {
		return &pb.GetUserByIDResponse{Status: nil, Data: nil}, nil
	}

	user, err := h.userSvc.GetUserByID(ctx, id)
	if err != nil {
		return &pb.GetUserByIDResponse{Status: nil, Data: nil}, nil
	}

	resp := &pb.GetUserByIDResponse{
		Status: &pb.PStatus{Success: true},
		Data: &pb.UserComponent{
			Id:           user.ID,
			Email:        *user.Email,
			Username:     user.Username,
			DisplayName:  user.DisplayName,
		},
	}

	return resp, nil
}