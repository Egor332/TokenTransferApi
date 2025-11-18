package testutil

import (
	"sync"
	"testing"

	"github.com/Egor332/TokenTransferApi/database"
	"github.com/Egor332/TokenTransferApi/pkg/common"
	"github.com/Egor332/TokenTransferApi/repository"
	"github.com/Egor332/TokenTransferApi/service"
	"github.com/stretchr/testify/require"
)

func TestTransfer_RaceCondition(t *testing.T) {
	database.Connect()

	repo := repository.NewGormWalletRepository()
	svc := service.NewWalletTransferService(repo, database.DB)

	var initialBalance int64 = 10

	subject := createOneWallet(t, database.DB, initialBalance, "A")
	receiver1 := createOneWallet(t, database.DB, 0, "B")
	receiver2 := createOneWallet(t, database.DB, 0, "C")
	sender1 := createOneWallet(t, database.DB, 100, "D")

	var wg sync.WaitGroup

	errChan := make(chan error, 3)

	wg.Add(3)

	// Routine A: Transfer 4 OUT (Subject -> Receiver1)
	go func() {
		defer wg.Done()
		_, err := svc.Transfer(subject, receiver1, 4)
		errChan <- err
	}()

	// Routine B: Transfer 7 OUT (Subject -> Receiver2)
	go func() {
		defer wg.Done()
		_, err := svc.Transfer(subject, receiver2, 7)
		errChan <- err
	}()

	// Routine C: Transfer 1 IN (Sender1 -> Subject)
	go func() {
		defer wg.Done()
		_, err := svc.Transfer(sender1, subject, 1)
		errChan <- err
	}()

	wg.Wait()
	close(errChan)

	var successCount int
	var insufficientFundsCount int

	for err := range errChan {
		if err == nil {
			successCount++
		} else if err == common.ErrInsufficientFunds {
			insufficientFundsCount++
		} else {
			require.NoError(t, err, "Unexpected error type during concurrency test")
		}
	}

	updatedSubject, err := repo.GetWalletByAddress(database.DB, subject)
	require.NoError(t, err)

	finalBalance := updatedSubject.Balance

	switch finalBalance {
	case 0:
		// Means (+1) happened early enough to facilitate both withdrawals
		require.Equal(t, 3, successCount, "If balance is 0, all txs should have succeeded")
	case 4:
		// Means (-7) succeeded, (-4) failed, (+1) succeeded
		require.Equal(t, 2, successCount)
		require.Equal(t, 1, insufficientFundsCount)
	case 7:
		// Means (-4) succeeded, (-7) failed, (+1) succeeded
		require.Equal(t, 2, successCount)
		require.Equal(t, 1, insufficientFundsCount)
	default:
		require.Fail(t, "Invalid final balance", "Balance %d is mathematically impossible if locking works correctly", finalBalance)
	}

	t.Logf("Concurrency Test Passed. Final Balance: %d, Successes: %d", finalBalance, successCount)
}
