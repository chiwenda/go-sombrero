/*
 * @Description: 配置gorm
 * @Author: chiwenda
 * @Date: 2022-10-23 16:41:17
 * @LastEditTime: 2022-10-23 18:25:19
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/config/gorm.go
 */
package config

import (
	"g-sombrero/context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	url := context.CTX_CONFIG.Mysql.Url
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to mysql fail:%v", nil)
	}
	mdb, err := db.DB()
	if err != nil {
		log.Fatalf("set db pool fail:%v", nil)
	}
	mdb.SetMaxIdleConns(10)  //最大空闲连接数
	mdb.SetMaxOpenConns(100) //打开数据库连接的最大数量
	mdb.SetConnMaxLifetime(time.Hour)
	return db
}
