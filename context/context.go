/*
 * @Description:
 * @Author: chiwenda
 * @Date: 2022-09-25 22:32:08
 * @LastEditTime: 2022-10-15 20:59:25
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/context/context.go
 */
package context

import (
	"g-sombrero/config/entity"

	"github.com/spf13/viper"
)

var (
	CTX_VIPER  *viper.Viper //viper
	CTX_CONFIG *entity.Config
)
