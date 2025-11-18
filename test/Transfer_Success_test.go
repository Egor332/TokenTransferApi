package testutil

import (
	"testing"

	"github.com/Egor332/TokenTransferApi/database"
	"github.com/Egor332/TokenTransferApi/repository"
	"github.com/Egor332/TokenTransferApi/service"
	"github.com/stretchr/testify/require"
)

func TestTransferSuccess(t *testing.T) {
	database.Connect()

	repo := repository.NewGormWalletRepository()
	svc := service.NewWalletTransferService(repo, database.DB)

	var fromStartBalance int64 = 5
	var toStartBalance int64 = 5
	var transferAmount int64 = 2

	from, to := createWallets(t, database.DB, fromStartBalance, toStartBalance)

	newBal, err := svc.Transfer(from, to, transferAmount)
	deleteWallets(t, database.DB)

	require.NoError(t, err)
	require.Equal(t, fromStartBalance-transferAmount, newBal)
}
