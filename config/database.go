package config

// Database 数据库相关配置
type Database struct {
	Domain   string // Domain 数据库服务器IP地址
	Port     string // Port 数据库端口
	User     string // User 用户名
	Password string // Password 密码
	Name     string // Name 数据库名
}

// GenConnArgs 根据配置拼接连接数据库的必要信息
func (d *Database) GenConnArgs() string {
	return d.User + ":" + d.Password + "@tcp(" + d.Domain + ":" + d.Port + ")/" + d.Name + "?charset=utf8&parseTime=True&loc=Local"
}
