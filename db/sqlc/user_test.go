package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createUserFunction(t)
}

func createUserFunction(t *testing.T) User {
	ars := CreateUserParams{
		Username:       utils.RandomString(3),
		HashedPassword: "secret",
		FullName:       utils.RandomString(3),
		Email:          fmt.Sprintf("%s@gmail.com", utils.RandomString(8)),
	}

	user, err := testQueries.CreateUser(context.Background(), ars)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user

}
