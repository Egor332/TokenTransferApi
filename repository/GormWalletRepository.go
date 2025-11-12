package repository

import (
	"errors"

	"github.com/Egor332/TokenTransferApi/models"
	"gorm.io/gorm"
)

type GormWalletRepository struct {
	DB *gorm.DB
}

var _ WalletRepositoryInterface = (*GormWalletRepository)(nil)

func NewGormWalletRepository(db *gorm.DB) *GormWalletRepository {
	return &GormWalletRepository{DB: db}
}

func (r *GormWalletRepository) AddToBalance(address string, additionAmount int64) error {
	updateResult := r.DB.Model(&models.Wallet{}).
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

func (r *GormWalletRepository) AddWallet(wallet *models.Wallet) error {
	result := r.DB.Create(wallet)
	return result.Error
}

func (r *GormWalletRepository) GetWalletByAddress(address string) (*models.Wallet, error) {
	var wallet models.Wallet
	result := r.DB.First(&wallet, "wallet_address=?", address)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &wallet, result.Error
}
