package repository

import (
	"errors"

	"github.com/Egor332/TokenTransferApi/models"
	"github.com/Egor332/TokenTransferApi/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormWalletRepository struct {
}

var _ WalletRepositoryInterface = (*GormWalletRepository)(nil)

func NewGormWalletRepository() *GormWalletRepository {
	return &GormWalletRepository{}
}

func (r *GormWalletRepository) SetNewBalance(db *gorm.DB, address string, newBalance int64) error {
	updateResult := db.Model(&models.Wallet{}).
		Where("wallet_address = ?", address).
		Update("balance", newBalance)

	if updateResult.Error != nil {
		return updateResult.Error
	}

	if updateResult.RowsAffected == 0 {
		return common.ErrWalletNotFound
	}

	return nil
}

func (r *GormWalletRepository) AddWallet(db *gorm.DB, wallet *models.Wallet) error {
	result := db.Create(wallet)
	return result.Error
}

func (r *GormWalletRepository) GetWalletByAddress(db *gorm.DB, address string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := db.Where("wallet_address = ?", address).First(&wallet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrWalletNotFound
		}
		return nil, err
	}

	return &wallet, nil
}

func (r *GormWalletRepository) GetWalletByAddressWithLock(db *gorm.DB, address string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("wallet_address = ?", address).First(&wallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrWalletNotFound
		}
		return nil, err
	}
	return &wallet, nil
}
