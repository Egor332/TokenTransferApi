package service

import (
	"github.com/Egor332/TokenTransferApi/models"
	"github.com/Egor332/TokenTransferApi/pkg/common"
	"github.com/Egor332/TokenTransferApi/repository"
	"gorm.io/gorm"
)

type WalletTransferService struct {
	repo repository.WalletRepositoryInterface
	db   *gorm.DB
}

func NewWalletTransferService(repo repository.WalletRepositoryInterface, db *gorm.DB) *WalletTransferService {
	return &WalletTransferService{
		repo: repo,
		db:   db,
	}
}

func (s *WalletTransferService) Transfer(fromAddress string, toAddress string, amount int64) (int64, error) {
	if amount <= 0 {
		return 0, common.ErrInvalidAmount
	}

	if fromAddress == toAddress {
		return 0, common.ErrSameWallet
	}

	var newFromBalance int64

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var fromWallet, toWallet *models.Wallet
		var err error

		// Deterministic Ordering strategy to escape deadlock
		if fromAddress < toAddress {
			fromWallet, err = s.repo.GetWalletByAddressWithLock(tx, fromAddress)
			if err != nil {
				return err
			}

			toWallet, err = s.repo.GetWalletByAddressWithLock(tx, toAddress)
			if err != nil {
				return err
			}
		} else {
			toWallet, err = s.repo.GetWalletByAddressWithLock(tx, toAddress)
			if err != nil {
				return err
			}

			fromWallet, err = s.repo.GetWalletByAddressWithLock(tx, fromAddress)
			if err != nil {
				return err
			}
		}

		if fromWallet.Balance < amount {
			return common.ErrInsufficientFunds
		}

		newFromBalance = fromWallet.Balance - amount
		err = s.repo.SetNewBalance(tx, fromAddress, newFromBalance)
		if err != nil {
			return err
		}

		newToBalance := toWallet.Balance + amount
		err = s.repo.SetNewBalance(tx, toAddress, newToBalance)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return newFromBalance, nil
}
