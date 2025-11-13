package testutil

import (
	"testing"

	"github.com/Egor332/TokenTransferApi/repository"
	"github.com/Egor332/TokenTransferApi/service"
	"github.com/stretchr/testify/require"
)

func TestTransfer_Success(t *testing.T) {
	db, cleanup := SetupTestDB(t)
	defer cleanup()

	repo := repository.NewGormWalletRepository()
	svc := service.NewWalletTransferService(repo, db)

	var fromStartBalance int64 = 5
	var toStartBalance int64 = 5
	var transferAmount int64 = 2

	from, to := setupWallets(t, db, fromStartBalance, toStartBalance)

	newBal, err := svc.Transfer(from, to, transferAmount)

	require.NoError(t, err)
	require.Equal(t, fromStartBalance-transferAmount, newBal)
}
