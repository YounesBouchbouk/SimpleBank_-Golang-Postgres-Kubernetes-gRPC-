package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/api"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/gapi"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	//to run GIN server
	// runGinServer(config, store)

	//run gRpc server
	rungRpcServer(config, store)

}

func rungRpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSimpleBankServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)

	if err != nil {
		log.Fatal("cannot create listner", err)
	}

	log.Printf("start gRpc server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
