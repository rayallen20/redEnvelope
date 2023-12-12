package config

import (
	"github.com/spf13/viper"
	"os"
	"redEnvelope/lib"
	"strings"
)

// env 用于区分生产/开发环境的变量 默认为开发环境
// dev: 开发环境
// proc: 生产环境
var env = "dev"

const (
	// configFileName 配置文件名
	configFileName = "config"
)

// Conf 配置对象
type Conf struct {
	Server   Server   // Server web服务相关配置
	Database Database // Database 数据库相关配置
}

// Load 读取配置文件至Config对象
func (c *Conf) Load() {
	configFilePath := getConfigFilePath()
	checkEnv()
	cfgFileName := configFileName + "." + env

	v := viper.New()
	v.AddConfigPath(configFilePath)
	v.SetConfigName(cfgFileName)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("load config failed:" + err.Error())
	}

	err = v.Unmarshal(c)
	if err != nil {
		panic("unmarshal config failed:" + err.Error())
	}
}

// getConfigFilePath 获取默认配置文件的路径
func getConfigFilePath() string {
	currentPath := lib.GetCurrentPath()
	currentPathSegments := strings.Split(currentPath, "/")
	currentPathSegments = currentPathSegments[:len(currentPathSegments)-2]
	currentPath = strings.Join(currentPathSegments, "/")
	currentPath += "/"
	return currentPath
}

// checkEnv 根据环境变量设置包内变量 env 的值
func checkEnv() {
	envValue := os.Getenv("ENV")
	if envValue == "prod" {
		env = "prod"
	}
}
