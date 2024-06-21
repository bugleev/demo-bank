package db

import (
	"context"
	"testing"

	"github.com/bugleev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, accountId int64) Entry {
	arg := CreateEntryParams{
		AccountID: accountId,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, accountId, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	testAccount := createRandomAccount(t)
	createRandomEntry(t, testAccount.ID)
}

func TestGetEntry(t *testing.T) {
	testAccount := createRandomAccount(t)
	testEntry := createRandomEntry(t, testAccount.ID)
	entryInDb, err := testQueries.GetEntry(context.Background(), testEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entryInDb)

	require.Equal(t, entryInDb.ID, testEntry.ID)
	require.Equal(t, entryInDb.AccountID, testEntry.AccountID)
	require.Equal(t, entryInDb.Amount, testEntry.Amount)
	require.Equal(t, entryInDb.CreatedAt, testEntry.CreatedAt)
}

func TestListEntries(t *testing.T) {
	testAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, testAccount.ID)
	}

	arg := ListEntriesParams{
		AccountID: testAccount.ID,
		Limit:     5,
		Offset:    5,
	}

	entriesList, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entriesList, 5)

	for _, entry := range entriesList {
		require.NotEmpty(t, entry)
	}
}
