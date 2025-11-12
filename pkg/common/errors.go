package common

import "errors"

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrWalletNotFound    = errors.New("wallet not found")
	ErrSameWallet        = errors.New("cannot transfer to the same wallet")
	ErrInvalidAmount     = errors.New("transfer amount must be positive")
)
