/*
 * @Description:
 * @Author: chiwenda
 * @Date: 2022-09-25 22:32:08
 * @LastEditTime: 2022-10-23 20:11:34
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/context/context.go
 */
package context

import (
	"g-sombrero/config/entity"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CTX_VIPER  *viper.Viper //viper
	CTX_CONFIG *entity.Config
	CTX_DB     *gorm.DB    //gorm
	CTX_LOGGER *zap.Logger //zap
)
