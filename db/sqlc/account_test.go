package db

import (
	"context"
	"testing"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/utils"
	"github.com/stretchr/testify/require"
)

func createAccountfunction(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account

}

func TestCreateAccout(t *testing.T) {
	createAccountfunction(t)

}

func TestGetAccount(t *testing.T) {
	createAccount := createAccountfunction(t)

	account, err := testQueries.GetAccount(context.Background(), createAccount.ID)

	require.NoError(t, err)

	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	require.Equal(t, account.ID, createAccount.ID)

}

func TestDeleteAccout(t *testing.T) {
	createAccount := createAccountfunction(t)

	err := testQueries.DeleteAccount(context.Background(), createAccount.ID)

	require.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	createAccount := createAccountfunction(t)

	args := UpdateAccountParams{
		ID:      createAccount.ID,
		Balance: 2000,
	}
	account, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

}
