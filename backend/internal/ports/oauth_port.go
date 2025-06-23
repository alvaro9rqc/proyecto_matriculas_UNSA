package ports

import (
	"context"

	"github.com/enrollment/gen/db"
)

type OauthRepositoryInterface interface {
	GetAccountByEmail(ctx context.Context, email string) (db.Account, error)
	FullListAccounts(ctx context.Context) ([]db.Account, error)
	ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error)
	CreateAccount(ctx context.Context, arg db.CreateAccountParams) error
	CreateAccountSession(ctx context.Context, arg db.CreateAccountSessionParams) (db.AccountSession, error)
	CreateOauthProvider(ctx context.Context, name string) error
	CreateAccountWithProviderName(ctx context.Context, arg db.CreateAccountWithProviderNameParams) error
	DeleteAccountByToken(ctx context.Context, token string) error
}
