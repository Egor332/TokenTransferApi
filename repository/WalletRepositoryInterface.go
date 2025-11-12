package repository

import "github.com/Egor332/TokenTransferApi/models"

type WalletRepositoryInterface interface {
	GetWalletByAddress(address string) (*models.Wallet, error)

	AddWallet(wallet *models.Wallet) error

	AddToBalance(address string, additionAmount int64) error
}
