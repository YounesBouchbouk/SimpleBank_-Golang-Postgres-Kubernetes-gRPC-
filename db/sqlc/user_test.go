package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createUserFunction(t)
}

func createUserFunction(t *testing.T) User {
	hashedPassword, err := utils.HashPassword("younes")

	require.NoError(t, err)

	ars := CreateUserParams{
		Username:       utils.RandomString(3),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomString(3),
		Email:          fmt.Sprintf("%s@gmail.com", utils.RandomString(8)),
	}

	user, err := testQueries.CreateUser(context.Background(), ars)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user

}

func TestUpdateUser(t *testing.T) {
	user1 := createUserFunction(t)
	require.NotEmpty(t, user1)

	email := fmt.Sprintf("%s@gmail.com", utils.RandomString(8))

	args := UpdateUserParams{
		Username: user1.Username,
		Email: sql.NullString{
			String: email,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), args)

	require.Empty(t, err)
	require.Nil(t, err)

	require.NotEqual(t, user1.Email, updatedUser.Email)
	require.Equal(t, updatedUser.Email, email)

}

func TestGetUser(t *testing.T) {
	user1 := createUserFunction(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
