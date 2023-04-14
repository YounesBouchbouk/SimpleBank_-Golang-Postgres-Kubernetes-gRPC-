package api

import (
	"fmt"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/token"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	tokenMake token.Maker
	store     *db.Store
	router    *gin.Engine
	config    utils.Config
}

func NewServer(config utils.Config, store *db.Store) (*Server, error) {

	//use token.NewJWTMaker to work with jwt
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	//use token.NewJWTMaker to work with NewPasetoMaker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, tokenMake: tokenMaker}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//user routers

	router.POST("/users", server.CreateUser)

	//account routers

	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.getAllAccounts)
	router.DELETE("/account/:id", server.deleteAccount)

	//transfer routers

	router.POST("/transfert", server.createTransfer)

	//authentification routes
	router.POST("/login", server.login)

	server.router = router
	return server, nil
}

// start and runs http server at specific address.
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
