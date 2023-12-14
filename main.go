package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"redEnvelope/config"
	"redEnvelope/controller/v1/user"
	"redEnvelope/customValid"
	"redEnvelope/lib/db"
)

func main() {
	// 加载配置文件
	c := &config.Conf{}
	c.Load()

	// 初始化句柄并尝试连接数据库
	db.Init(c.Database)

	r := gin.Default()
	err := r.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		panic("set trusted proxies failed:" + err.Error())
	}

	// 注册自定义验证规则
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("assert engine to validate failed")
	}
	err = validate.RegisterValidation("validBigFloat", customValid.ValidBigFloat)
	if err != nil {
		panic("register custom validator failed:" + err.Error())
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/user/deposit", user.Deposit)
	}

	addr := c.Server.Address + ":" + c.Server.Port
	r.Run(addr)
}
