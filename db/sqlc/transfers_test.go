package db

import (
	"context"
	"testing"

	"github.com/bugleev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountId1 int64, accountId2 int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: accountId1,
		ToAccountID:   accountId2,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, accountId1, transfer.FromAccountID)
	require.Equal(t, accountId2, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	testAccount1 := createRandomAccount(t)
	testAccount2 := createRandomAccount(t)
	createRandomTransfer(t, testAccount1.ID, testAccount2.ID)
}

func TestGetTransfer(t *testing.T) {
	testAccount1 := createRandomAccount(t)
	testAccount2 := createRandomAccount(t)
	testTransfer := createRandomTransfer(t, testAccount1.ID, testAccount2.ID)

	transferInDb, err := testQueries.GetTransfer(context.Background(), testTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transferInDb)

	require.Equal(t, transferInDb.ID, testTransfer.ID)
	require.Equal(t, transferInDb.FromAccountID, testTransfer.FromAccountID)
	require.Equal(t, transferInDb.ToAccountID, testTransfer.ToAccountID)
	require.Equal(t, transferInDb.Amount, testTransfer.Amount)
	require.Equal(t, transferInDb.CreatedAt, testTransfer.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	testAccount1 := createRandomAccount(t)
	testAccount2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, testAccount1.ID, testAccount2.ID)
	}

	arg := ListTransfersParams{
		FromAccountID: testAccount1.ID,
		ToAccountID:   testAccount2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfersList, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfersList, 5)

	for _, transfer := range transfersList {
		require.NotEmpty(t, transfer)
	}
}
