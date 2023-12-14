package deposit

import "redEnvelope/biz/user"

type Response struct {
	User        UserResponse           `json:"user"`
	Transaction TransactionLogResponse `json:"transaction"`
}

type UserResponse struct {
	Id      int    `json:"id"`      // Id 用户在saba侧的id
	Balance string `json:"balance"` // Balance 用户当前余额
}

type TransactionLogResponse struct {
	Id     int    `json:"id"`     // Id 交易流水id
	Status string `json:"status"` // Status 交易状态
}

func (r *Response) Fill(user user.User, transactionLog user.TransactionLog) {
	r.User.fill(user)
	r.Transaction.fill(transactionLog)
}

func (u *UserResponse) fill(user user.User) {
	u.Id = user.Id
	u.Balance = user.Balance.Text('f', 18)
}

func (t *TransactionLogResponse) fill(transaction user.TransactionLog) {
	t.Id = transaction.Id
	t.Status = transaction.Status
}
