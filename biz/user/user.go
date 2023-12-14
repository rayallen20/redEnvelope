package user

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
	"redEnvelope/lib/db"
	"redEnvelope/model"
	"redEnvelope/sysError"
)

const (
	MaxBalance = 99999.999999999999999999 // MaxBalance 最大余额
)

// User 用户对象
type User struct {
	Id      int        // Id 用户在本系统内的id
	SabaId  int        // SabaId 用户在saba侧的id
	Balance *big.Float // Balance 用户余额
	Status  string     // Status 用户状态
}

// Deposit 用户充值
func (u *User) Deposit(sabaId int, name string, amount *big.Float, contractAddress, fromWalletAddress, hash string) (transactionLog *TransactionLog, err error) {
	// step1. 查找或创建用户
	orm := &model.User{}
	err = orm.FindBySabaId(sabaId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = orm.Create(sabaId, name)
			if err != nil {
				return nil, &sysError.DBError{Message: err.Error()}
			}
		}
	}
	u.FillByOrm(orm)

	// step2. 创建交易流水记录
	transactionLogBiz := &TransactionLog{}
	transactionLogBiz.Deposit(*u, amount, contractAddress, fromWalletAddress, hash)

	// step3. 创建事务 更新用户余额同时创建交易流水记录
	err = u.deposit(amount, transactionLogBiz)
	if err != nil {
		return nil, &sysError.TransactionError{Message: err.Error()}
	}

	return transactionLogBiz, nil
}

// deposit 更新用户余额同时创建交易流水记录(完成了修改数据库/落盘数据库的操作)
func (u *User) deposit(amount *big.Float, transactionLog *TransactionLog) (err error) {
	tx := db.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	// step1. 加行锁查询用户数据 并 更新用户余额
	userOrm := &model.User{}
	userOrm.Id = u.Id
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where(userOrm).First(userOrm).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	nowBalance := new(big.Float)
	nowBalance.SetString(userOrm.Balance)
	userOrm.Balance = new(big.Float).Add(nowBalance, amount).Text('f', 18)
	err = tx.Save(userOrm).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// step2. 创建交易流水记录
	transactionLogOrm := &model.TransactionLog{}
	transactionLogOrm.UserId = userOrm.Id
	transactionLogOrm.Type = transactionLog.Type
	transactionLogOrm.Status = transactionLog.Status
	transactionLogOrm.BeforeTransactBalance = transactionLog.BeforeTransactBalance.Text('f', 18)
	transactionLogOrm.TransactBalance = transactionLog.TransactBalance.Text('f', 18)
	transactionLogOrm.AfterTransactBalance = transactionLog.AfterTransactBalance.Text('f', 18)
	transactionLogOrm.FromWalletAddress = transactionLog.FromWalletAddress
	transactionLogOrm.ToWalletAddress = transactionLog.ToWalletAddress
	transactionLogOrm.TransactionHash = transactionLog.TransactionHash
	transactionLogOrm.ContractAddress = transactionLog.ContractAddress
	err = tx.Save(transactionLogOrm).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// step3. 事务执行成功后更新用户对象与交易流水对象id
	u.FillByOrm(userOrm)
	transactionLog.Id = transactionLogOrm.Id

	return nil
}

// FillByOrm 根据user表orm对象填充用户对象
func (u *User) FillByOrm(orm *model.User) {
	u.Id = orm.Id
	u.SabaId = orm.SabaId
	u.Status = orm.Status
	u.Balance = new(big.Float)
	// Tips: 这里我认为只要是能写到DB中的big.Float 都是合法的
	// Tips: 所以在读的时候不需要判断是否合法
	u.Balance, _ = u.Balance.SetString(orm.Balance)
}
