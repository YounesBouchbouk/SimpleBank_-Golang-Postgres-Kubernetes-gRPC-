package gapi

import (
	"context"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/pb"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't hash password")
	}

	args := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetPassword(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, args)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can't create user")

	}

	rsp := &pb.CreateUserResponse{
		User: transferFromToUserResp(user),
	}

	return rsp, nil

}
