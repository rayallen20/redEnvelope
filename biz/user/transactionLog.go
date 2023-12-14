package user

import (
	"math/big"
	"redEnvelope/config"
)

const (
	TransactionTypeDeposit = "DEPOSIT" // TransactionTypeDeposit 充值
)

const (
	TransactionStatusComplete = "COMPLETE" // TransactionStatusComplete 交易完成
	TransactionStatusPending  = "PENDING"  // TransactionStatusPending 交易进行中
	TransactionStatusApprove  = "APPROVE"  // TransactionStatusApprove 审核通过
	TransactionStatusRefuse   = "REFUSE"   // TransactionStatusRefuse 审核拒绝
)

// TransactionLog 交易流水记录
type TransactionLog struct {
	Id                    int        // Id 交易流水id
	User                  User       // User 交易发起者
	Type                  string     // Type 交易类型
	Status                string     // Status 交易状态
	BeforeTransactBalance *big.Float // BeforeTransactBalance 交易前余额
	TransactBalance       *big.Float // TransactBalance 当次交易金额 (正数表示收入,负数表示支出)
	AfterTransactBalance  *big.Float // AfterTransactBalance 交易后余额
	FromWalletAddress     string     // FromWalletAddress 交易发起的钱包地址
	ToWalletAddress       string     // ToWalletAddress 交易接收的钱包地址
	TransactionHash       string     // TransactionHash 交易hash
	ContractAddress       string     // ContractAddress 交易合约地址
}

// Deposit 创建用户充值交易流水对象
func (t *TransactionLog) Deposit(user User, amount *big.Float, contractAddress, fromWalletAddress, hash string) {
	t.User = user
	t.Type = TransactionTypeDeposit
	t.Status = TransactionStatusComplete
	t.BeforeTransactBalance = user.Balance
	t.TransactBalance = amount
	t.ContractAddress = contractAddress
	t.FromWalletAddress = fromWalletAddress
	t.ToWalletAddress = config.ConfigObj.CenterWalletAddress
	t.TransactionHash = hash
	t.AfterTransactBalance = new(big.Float).Add(user.Balance, amount)
}
