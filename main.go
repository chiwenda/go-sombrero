/*
 * @Description:
 * @Author: chiwenda
 * @Date: 2022-09-23 21:05:00
 * @LastEditTime: 2022-10-23 20:11:55
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/main.go
 */
package main

import (
	"fmt"
	"g-sombrero/config"
	ctx "g-sombrero/context"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//初始化viper
	ctx.CTX_VIPER = config.Viper()
	//初始化数据库
	ctx.CTX_DB = config.Gorm()
	//初始化zap日志
	ctx.CTX_LOGGER = config.Zap()
	//启动fiber服务
	app := fiber.New(fiber.Config{
		AppName: ctx.CTX_CONFIG.Name,
	})
	app.Listen(":" + fmt.Sprint(ctx.CTX_CONFIG.Port))
}
