package models

type Transaction struct {
	ID          int64  `gorm:"column:id;PrimaryKey"`
	FromAddress string `gorm:"column:from_address;type:varchar(42);not null"`
	ToAddress   string `gorm:"column:to_address;type:varchar(42);not null"`
	Amount      int64  `gorm:"column:amount;not null"`
}
