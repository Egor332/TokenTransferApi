package repository

import (
	"github.com/Egor332/TokenTransferApi/models"
	"gorm.io/gorm"
)

type WalletRepositoryInterface interface {
	GetWalletByAddress(db *gorm.DB, address string) (*models.Wallet, error)

	AddWallet(db *gorm.DB, wallet *models.Wallet) error

	AddToBalance(db *gorm.DB, address string, additionAmount int64) error

	GetWalletByAddressWithLock(db *gorm.DB, address string) (*models.Wallet, error)
}
