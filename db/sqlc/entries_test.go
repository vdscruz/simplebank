package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vdscruz/simplebank/util"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestListEntries(t *testing.T) {

	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createEntryFromAccount(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     20,
		Offset:    0,
	}

	entries, err := testStore.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 10)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, account.ID)
	}
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testStore.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func createEntryFromAccount(t *testing.T, account Account) Entry {

	args_entry := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testStore.CreateEntry(context.Background(), args_entry)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args_entry.AccountID, entry.AccountID)
	require.Equal(t, args_entry.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func createRandomEntry(t *testing.T) Entry {
	arg_account := GetAccountsParams{
		Limit:  20,
		Offset: 0,
	}

	accounts, err := testStore.GetAccounts(context.Background(), arg_account)
	require.NoError(t, err)

	index := util.RandomInt(0, int64(len(accounts)-1))
	args_entry := CreateEntryParams{
		AccountID: accounts[index].ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testStore.CreateEntry(context.Background(), args_entry)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args_entry.AccountID, entry.AccountID)
	require.Equal(t, args_entry.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
