package api

import (
	"database/sql"
	"net/http"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	args := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(ctx, args)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountRequest struct {
	ID int64 `uri:"id"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	account, err := server.store.GetAccount(ctx, req.ID)

	if account.Owner != authPayload.Username {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "you don't have permision to acces this account"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type getAllAccountsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (server *Server) getAllAccounts(ctx *gin.Context) {
	var req getAllAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountsParams{
		Offsetnb: (req.PageID - 1) * req.PageSize,
		Limitnb:  int32(req.PageSize),
		Username: authPayload.Username,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)

}

type deleteAccountRequest struct {
	ID int64 `uri:"id"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	println(req.ID)

	account, err := server.store.GetAccount(ctx, req.ID)

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if account.Owner != authPayload.Username {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "you don't have permition to do this"})
		return
	}

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user doesn't exist"})
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteAccount(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
