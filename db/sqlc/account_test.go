package db

import (
	"context"
	"testing"

	"github.com/YounesBouchbouk/SimpleBank_-Golang-Postgres-Kubernetes-gRPC-/db/utils"
	"github.com/stretchr/testify/require"
)

func createAccountfunction(t *testing.T) Account {
	arg := accountsParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.accounts(context.Background(), arg)

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

	account, err := testQueries.GetAccounts(context.Background(), createAccount.ID)

	require.NoError(t, err)

	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	require.Equal(t, account.ID, createAccount.ID)

}

func TestDeleteAccout(t *testing.T) {
	createAccount := createAccountfunction(t)

	err := testQueries.DeleteAccounts(context.Background(), createAccount.ID)

	require.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	createAccount := createAccountfunction(t)

	args := UpdateAuthorParams{
		ID:      createAccount.ID,
		Balance: 2000,
	}
	err := testQueries.UpdateAuthor(context.Background(), args)

	require.NoError(t, err)

}
