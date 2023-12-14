package model

import (
	"redEnvelope/lib/db"
	"time"
)

const (
	ActiveStatus = "ACTIVE" // ActiveStatus 活跃状态
	BannedStatus = "BANNED" // BannedStatus 禁用状态
)

type User struct {
	Id        int       `gorm:"primaryKey"` // Id 主键自增Id
	SabaId    int       // SabaId 用户在saba侧的id
	Name      string    // Name 用户名
	Balance   string    // Balance 用户余额
	Status    string    // Status 用户状态
	CreatedAt time.Time `gorm:"autoCreateTime"` // CreatedAt 数据创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // UpdatedAt 数据更新时间
}

func (u *User) FindBySabaId(sabaId int) (err error) {
	u.SabaId = sabaId
	result := db.Conn.Where(u).First(u)
	return result.Error
}

func (u *User) FindById(id int) (err error) {
	u.Id = id
	result := db.Conn.Where(u).First(u)
	return result.Error
}

func (u *User) Create(sabaId int, name string) (err error) {
	u.SabaId = sabaId
	u.Name = name
	u.Status = ActiveStatus
	u.Balance = "0"
	return db.Conn.Create(u).Error
}
