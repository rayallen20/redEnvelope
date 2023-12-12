package main

import (
	"github.com/gin-gonic/gin"
	"redEnvelope/config"
	"redEnvelope/lib/db"
)

func main() {
	c := &config.Conf{}
	c.Load()

	db.Init(c.Database)

	r := gin.Default()
	err := r.SetTrustedProxies([]string{"http", "https"})
	if err != nil {
		panic("set trusted proxies failed:" + err.Error())
	}

	r.Run("0.0.0.0:8094")
}
