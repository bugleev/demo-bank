package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bugleev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandmCurrency(),
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
	testAccount := createRandomAccount(t)
	accountInDb, err := testQueries.GetAccount(context.Background(), testAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accountInDb)

	require.Equal(t, accountInDb.ID, testAccount.ID)
	require.Equal(t, accountInDb.Owner, testAccount.Owner)
	require.Equal(t, accountInDb.Balance, testAccount.Balance)
	require.Equal(t, accountInDb.Currency, testAccount.Currency)
	require.Equal(t, accountInDb.CreatedAt, testAccount.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	testAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      testAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, updatedAccount.ID, testAccount.ID)
	require.Equal(t, updatedAccount.Owner, testAccount.Owner)
	require.Equal(t, updatedAccount.Balance, arg.Balance)
	require.Equal(t, updatedAccount.Currency, testAccount.Currency)
	require.Equal(t, updatedAccount.CreatedAt, testAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	testAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), testAccount.ID)

	require.NoError(t, err)

	accountInDb, err := testQueries.GetAccount(context.Background(), testAccount.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountInDb)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accountsList, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accountsList, 5)

	for _, account := range accountsList {
		require.NotEmpty(t, account)
	}
}
