package api

import (
	"net/http"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username       string `json:"username" binding:"required"`
	HashedPassword string `json:"password" binding:"required"`
	FullName       string `json:"full_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.HashedPassword)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}

	rsp := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}
