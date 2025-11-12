package models

type Wallet struct {
	WalletAdress string `gorm:"column:wallet_address;primaryKey;type:varchar(42)"`
	Balance      int64  `gorm:"column:balance;not null"`
}
