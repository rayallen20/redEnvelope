package lib

import (
	"runtime"
)

// GetCurrentPath 获取当前路径
func GetCurrentPath() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic("Call runtime.Caller() failed")
	}

	return path
}
