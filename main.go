package main

import (
	"database/sql"
	"log"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/api"
	db "github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgres://younes:secret@localhost:5432/simple_bank?sslmode=disable"
	APIaddress = "0.0.0.0:8080"
)

func main() {
	var err error
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connection)

	server := api.NewServer(store)

	err = server.Start(APIaddress)

	if err != nil {
		log.Fatal("cannot start the server:", err)

	}

}
