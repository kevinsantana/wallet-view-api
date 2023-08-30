package wallet

import (
	"context"

	"github.com/kevinsantana/wallet-view-api/internal/share"
	"github.com/kevinsantana/wallet-view-api/internal/config"
)

type WalletView interface {
	Balance(ctx context.Context, hexAddress string, currency string) (share.Balance, error)
}

type WalletViewUseCase struct {
	cfg *config.Config
}

func NewWalletViewUseCase (cfg *config.Config) WalletViewUseCase {
	return WalletViewUseCase {
		cfg: cfg,
	}
}

func (w WalletViewUseCase) Balance(ctx context.Context, hexAddress string, currency string) (share.Balance, error) {
	return share.Balance{}, nil
}