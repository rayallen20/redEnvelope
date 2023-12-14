package model

import "time"

type TransactionLog struct {
	Id                       int       `gorm:"primaryKey"` // Id 交易流水id
	UserId                   int       // UserId 交易发起者
	Type                     string    // Type 交易类型
	Status                   string    // Status 交易状态
	User                     *User     `gorm:"foreignKey:UserId"` // User 交易发起者
	WithdrawTransactionLogId int       // WithdrawTransactionLogId 提现交易流水id TODO: 本字段未来会被废弃
	RedEnvelopeId            int       // RedEnvelopeId 红包id
	RedEnvelopeItemId        int       // RedEnvelopeItemId 红包条目id
	TransferId               int       // TransferId 转账id
	BeforeTransactBalance    string    // BeforeTransactBalance 交易前余额
	TransactBalance          string    // TransactBalance 当次交易金额 (正数表示收入,负数表示支出)
	AfterTransactBalance     string    // AfterTransactBalance 交易后余额
	FromWalletAddress        string    // FromWalletAddress 交易发起的钱包地址
	ToWalletAddress          string    // ToWalletAddress 交易接收的钱包地址
	TransactionHash          string    // TransactionHash 交易hash
	ContractAddress          string    // ContractAddress 交易合约地址
	CreatedAt                time.Time `gorm:"autoCreateTime"` // CreatedAt 数据创建时间
	UpdatedAt                time.Time `gorm:"autoUpdateTime"` // UpdatedAt 数据更新时间
}
