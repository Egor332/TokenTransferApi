package testutil

import (
	"testing"

	"github.com/Egor332/TokenTransferApi/models"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createWallets(t *testing.T, db *gorm.DB, balance1 int64, balance2 int64) (string, string) {
	w1 := models.Wallet{WalletAdress: "0x0000000000000000000000000000000000000000", Balance: balance1}
	w2 := models.Wallet{WalletAdress: "0x0000000000000000000000000000000000000001", Balance: balance2}

	require.NoError(t, db.Create(&w1).Error)
	require.NoError(t, db.Create(&w2).Error)

	return w1.WalletAdress, w2.WalletAdress
}

func createOneWallet(t *testing.T, db *gorm.DB, balance int64, address string) string {
	w := models.Wallet{WalletAdress: address, Balance: balance}

	require.NoError(t, db.Create(&w).Error)

	return address
}

func deleteWallets(t *testing.T, db *gorm.DB) {
	db.Where("1 = 1").Delete(&models.Wallet{})
}
