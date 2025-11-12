package repository

import (
	"errors"

	"github.com/Egor332/TokenTransferApi/models"
	"gorm.io/gorm"
)

type GormWalletRepository struct {
}

var _ WalletRepositoryInterface = (*GormWalletRepository)(nil)

func NewGormWalletRepository() *GormWalletRepository {
	return &GormWalletRepository{}
}

func (r *GormWalletRepository) AddToBalance(db *gorm.DB, address string, additionAmount int64) error {
	updateResult := db.Model(&models.Wallet{}).
		Where("wallet_address = ?", address).
		Update("balance", gorm.Expr("balance + ?", additionAmount))

	if updateResult.Error != nil {
		return updateResult.Error
	}

	if updateResult.RowsAffected == 0 {
		return errors.New("wallet update failed: no rows affected")
	}

	return nil
}

func (r *GormWalletRepository) AddWallet(db *gorm.DB, wallet *models.Wallet) error {
	result := db.Create(wallet)
	return result.Error
}

func (r *GormWalletRepository) GetWalletByAddress(db *gorm.DB, address string) (*models.Wallet, error) {
	var wallet models.Wallet
	result := db.First(&wallet, "wallet_address=?", address)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &wallet, result.Error
}
