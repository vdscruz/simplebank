package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	ListTransfersByDestiny(ctx context.Context, accountID int64) ([]Transfer, error)
	ListTransfersByOrigin(ctx context.Context, accountID int64) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
