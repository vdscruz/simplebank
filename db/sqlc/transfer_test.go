package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vdscruz/simplebank/util"
)

func TestCreateTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)

	createRandomTransfer(t, from, to)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t, createRandomAccount(t), createRandomAccount(t))
	transfer2, err := testStore.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer.ID, transfer2.ID)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from, to)
	}

	arg := ListTransfersParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Limit:         20,
		Offset:        0,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 10)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, from.ID, transfer.FromAccountID)
		require.Equal(t, to.ID, transfer.ToAccountID)
		require.NotEmpty(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
	}
}

func TestListTransfersByDestiny(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from, to)
	}

	transfers, err := testStore.ListTransfersByDestiny(context.Background(), to.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, to.ID, transfer.ToAccountID)
	}
}

func TestListTransfersByOrigin(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from, to)
	}

	transfers, err := testStore.ListTransfersByOrigin(context.Background(), from.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, from.ID, transfer.FromAccountID)
	}
}

func createRandomTransfer(t *testing.T, from Account, to Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, from.ID)
	require.Equal(t, transfer.ToAccountID, to.ID)
	require.Equal(t, transfer.Amount, arg.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
