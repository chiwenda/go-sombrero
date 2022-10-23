/*
 * @Description:
 * @Author: chiwenda
 * @Date: 2022-09-24 22:51:21
 * @LastEditTime: 2022-10-23 20:09:41
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/config/zap.go
 */
package config

import (
	"log"

	"go.uber.org/zap"
)

func Zap() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("init zap logger err:%v", err)
	}
	return logger
}
