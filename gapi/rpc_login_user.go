package gapi

import (
	"context"
	"database/sql"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/pb"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user doesn't exist")

		}
		return nil, status.Errorf(codes.Internal, err.Error())

	}

	err = utils.CheckPassword(user.HashedPassword, req.Password)

	if err != nil {

		return nil, status.Errorf(codes.NotFound, "wrong password")

	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create accessToken")

	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenduration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create refresh Token")

	}

	createSessionPayload := db.CreateSessionParams{
		ID:           refreshTokenPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    server.extractMetadata(ctx).UserAgent,
		ClientIp:     server.extractMetadata(ctx).ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
	}

	session, err := server.store.CreateSession(ctx, createSessionPayload)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't create session")

	}

	res := &pb.LoginUserResponse{
		User:                  transferFromToUserResp(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessTokenPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(accessTokenPayload.ExpiredAt),
	}

	return res, nil
}
