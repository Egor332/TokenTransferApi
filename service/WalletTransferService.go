package service

import (
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

func (s *WalletTransferService) Transfer(fromAddress string, toAddress string, amount int64) error {
	if amount <= 0 {
		return common.ErrInvalidAmount
	}

	if fromAddress == toAddress {
		return common.ErrSameWallet
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		fromWallet, err := s.repo.GetWalletByAddressWithLock(tx, fromAddress)
		if err != nil {
			return err
		}
		if fromWallet.Balance < amount {
			return common.ErrInsufficientFunds
		}

		toWallet, err := s.repo.GetWalletByAddressWithLock(tx, toAddress)
		if err != nil {
			return err
		}

		newFromBalance := fromWallet.Balance - amount
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
		return err
	}

	return nil
}
