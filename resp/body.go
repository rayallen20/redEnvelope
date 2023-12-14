package resp

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"redEnvelope/sysError"
)

type Body struct {
	Code    int    `json:"code"`    // Code 响应码
	Message string `json:"message"` // Message 响应信息
	Data    any    `json:"data"`    // Data 响应载荷
}

// 错误码规则:
// 101XX:客户端错误
// 102XX:服务端错误
// 103XX:业务逻辑错误

const (
	Success         = 200   // Success 成功响应
	JSONSyntaxError = 10101 // JSONSyntaxError JSON语法错误
	ValidationErr   = 10102 // ValidationErr 参数校验错误
	DbErr           = 10201 // DbErr 操作数据库错误
	TransactionErr  = 10202 // TransactionErr 事务执行错误
)

// JSONSyntaxErr JSON语法错误
func (b *Body) JSONSyntaxErr(err *json.SyntaxError) {
	b.Code = JSONSyntaxError
	b.Message = err.Error()
	b.Data = make(map[string]any)
}

// ValidationErr 参数校验错误
func (b *Body) ValidationErr(err *validator.ValidationErrors) {
	b.Code = ValidationErr
	b.Message = err.Error()
	b.Data = make(map[string]any)
}

// DbErr 操作数据库错误
func (b *Body) DbErr(err *sysError.DBError) {
	b.Code = DbErr
	b.Message = err.Error()
	b.Data = make(map[string]any)
}

// TransactionErr 事务执行错误
func (b *Body) TransactionErr(err *sysError.TransactionError) {
	b.Code = TransactionErr
	b.Message = err.Error()
	b.Data = make(map[string]any)
}

// Success 成功响应
func (b *Body) Success(data any) {
	// 如果data为nil 则初始化data
	// 这样做的目的在于确保在JSON序列化后的data字段必定是一个JSON对象
	if data == nil {
		data = make(map[string]any)
	}
	b.Code = Success
	b.Message = "success"
	b.Data = data
}
