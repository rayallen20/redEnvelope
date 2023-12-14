package deposit

type Request struct {
	User        UserRequest        `json:"user" binding:"required"`        // User 用户对象
	Transaction TransactionRequest `json:"transaction" binding:"required"` // Transaction 交易对象
}

type UserRequest struct {
	Id   *int    `json:"id" binding:"required,gt=0"` // Id 用户在saba侧的id
	Name *string `json:"name" binding:"required"`    // Name 用户名
}

type TransactionRequest struct {
	ContractAddress *string `json:"contract_address" binding:"required"`     // ContractAddress 合约地址,标识币种
	WalletAddress   *string `json:"wallet_address" binding:"required"`       // WalletAddress 钱包地址
	Hash            *string `json:"hash" binding:"required"`                 // Hash 交易hash
	Amount          *string `json:"amount" binding:"required,validBigFloat"` // Amount 交易金额
}
