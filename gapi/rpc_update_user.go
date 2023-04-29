package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/pb"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	var ars db.UpdateUserParams

	if req.Email != "" {
		ars.Email = sql.NullString{
			String: req.Email,
			Valid:  true,
		}
	}

	if req.FullName != "" {
		ars.FullName = sql.NullString{
			String: req.FullName,
			Valid:  true,
		}
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)

		if err != nil {
			return nil, status.Errorf(codes.Internal, "can't hash password")

		}
		ars.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  true,
		}

		ars.PasswordChangedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	ars.Username = req.Username

	user, err := server.store.UpdateUser(ctx, ars)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't update user information")

	}

	res := &pb.UpdateUserResponse{
		User: transferFromToUserResp(user),
	}

	return res, nil

}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}
