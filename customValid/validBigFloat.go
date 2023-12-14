package customValid

import (
	"github.com/go-playground/validator/v10"
	"math/big"
)

// ValidBigFloat 自定义校验规则 判断一个字符串是否可以转换为big.Float
func ValidBigFloat(fl validator.FieldLevel) bool {
	bigFloat := new(big.Float)
	_, ok := bigFloat.SetString(fl.Field().String())
	return ok
}
