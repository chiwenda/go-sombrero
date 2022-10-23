/*
 * @Description: viper配置读取
 * @Author: chiwenda
 * @Date: 2022-10-15 16:34:05
 * @LastEditTime: 2022-10-15 20:29:40
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/core/viper.go
 */
package config

import (
	"fmt"
	"g-sombrero/context"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed :", e.Name)
		v.Unmarshal(&context.CTX_CONFIG)
	})
	v.WatchConfig()
	v.Unmarshal(&context.CTX_CONFIG)
	return v
}
