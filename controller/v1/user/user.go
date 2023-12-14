package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"math/big"
	"net/http"
	"redEnvelope/biz/user"
	"redEnvelope/param/user/deposit"
	"redEnvelope/resp"
	"redEnvelope/sysError"
)

// Deposit 用户充值
func Deposit(ctx *gin.Context) {
	// 获取参数
	reqParam := &deposit.Request{}
	respBody := &resp.Body{}
	err := ctx.ShouldBindJSON(reqParam)
	if err != nil {
		// 参数解析错误
		var jsonSyntaxErr *json.SyntaxError
		if errors.As(err, &jsonSyntaxErr) {
			respBody.JSONSyntaxErr(jsonSyntaxErr)
			ctx.JSON(http.StatusOK, respBody)
		}

		// 参数校验错误
		var vErr validator.ValidationErrors
		if errors.As(err, &vErr) {
			respBody.ValidationErr(&vErr)
			ctx.JSON(http.StatusOK, respBody)
			return
		}
	}
	// 业务逻辑
	amountBigFloat := new(big.Float)
	amountBigFloat, _ = amountBigFloat.SetString(*reqParam.Transaction.Amount)

	userBiz := &user.User{}
	transactionLogBiz, err := userBiz.Deposit(
		*reqParam.User.Id,
		*reqParam.User.Name,
		amountBigFloat,
		*reqParam.Transaction.ContractAddress,
		*reqParam.Transaction.WalletAddress,
		*reqParam.Transaction.Hash,
	)
	if err != nil {
		// 操作数据库错误
		var dbErr *sysError.DBError
		if errors.As(err, &dbErr) {
			respBody.DbErr(dbErr)
			ctx.JSON(http.StatusOK, respBody)
			return
		}

		// 事务执行错误
		var transactionErr *sysError.TransactionError
		if errors.As(err, &transactionErr) {
			respBody.TransactionErr(transactionErr)
			ctx.JSON(http.StatusOK, respBody)
			return
		}
	}

	// 返回响应
	response := &deposit.Response{}
	response.Fill(*userBiz, *transactionLogBiz)
	respBody.Success(response)
	ctx.JSON(http.StatusOK, respBody)
	return
}
