package testutil

import (
	"sync"
	"testing"

	"github.com/Egor332/TokenTransferApi/database"
	"github.com/Egor332/TokenTransferApi/repository"
	"github.com/Egor332/TokenTransferApi/service"
	"github.com/stretchr/testify/require"
)

func TestTransfer_RaceCondition(t *testing.T) {
	database.Connect()

	repo := repository.NewGormWalletRepository()
	svc := service.NewWalletTransferService(repo, database.DB)

	var fromStartBalance int64 = 5
	var toStartBalance int64 = 5
	var transferAmount int64 = 4

	from, to := createWallets(t, database.DB, fromStartBalance, toStartBalance)

	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		<-start
		_, _ = svc.Transfer(from, to, transferAmount)
	}()
	go func() {
		defer wg.Done()
		<-start
		_, _ = svc.Transfer(from, to, transferAmount)
	}()

	close(start)
	wg.Wait()

	var fromCurrentBalance int64
	var toCurrentBalance int64
	database.DB.Raw("SELECT balance FROM wallets WHERE wallet_address = ?", from).Scan(&fromCurrentBalance)
	database.DB.Raw("SELECT balance FROM wallets WHERE wallet_address = ?", to).Scan(&toCurrentBalance)
	deleteWallets(t, database.DB)
	require.Equal(t, fromStartBalance-transferAmount, fromCurrentBalance)
	require.Equal(t, toStartBalance+transferAmount, toCurrentBalance)
}
