package api

import (
	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
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
	return server
}

// start and runs http server at specific address.
func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
