package db

import (
	"bank/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	//create account
	expected := createRandomAccount(t)
	actual, err := testQueries.GetAccount(context.Background(), expected.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Balance, actual.Balance)
	require.Equal(t, expected.Currency, actual.Currency)
	require.Equal(t, expected.Owner, actual.Owner)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	expected := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      expected.ID,
		Balance: util.RandomMoney(),
	}

	actual, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, arg.Balance, actual.Balance)
	require.Equal(t, expected.Currency, actual.Currency)
	require.Equal(t, expected.Owner, actual.Owner)
	require.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	expected := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), expected.ID)
	require.NoError(t, err)

	acc, err := testQueries.GetAccount(context.Background(), expected.ID)
	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, acc)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
