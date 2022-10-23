package core

import (
	"flag"
	"fmt"
	"g-sombrero/core/common"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// Viper 优先级:命令行->
func Viper(path ...string) *viper.Viper {
	var config string //配置文件
	if len(path) == 0 {
		//从命令行获取
		flag.StringVar(&config, "conf", "", "choose config file")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(common.GlobalEnvName); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = common.DefaultConfigEnv
					fmt.Printf("the used gin mode is: %s env,config path is: %s\n", gin.DebugMode, common.DefaultConfigEnv)
				case gin.TestMode:
					config = common.TestConfigEnv
					fmt.Printf("the used gin mode is: %s env,config path is: %s\n", gin.TestMode, common.TestConfigEnv)
				case gin.ReleaseMode:
					config = common.ReleaseConfigEnv
					fmt.Printf("the used gin mode is: %s env,config path is: %s\n", gin.ReleaseMode, common.ReleaseConfigEnv)
				}
			} else {
				config = configEnv
				fmt.Printf("the used system mode is: %s env,config path is: %s\n", common.GlobalEnvName, config)
			}
		} else {
			fmt.Printf("the  cmd mode, command is: -config ,config path is: %s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("the used gin mode is new Viper() env,config path is: %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yml")
	//读取配置
	err := v.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed:%s \n", e.Name)
	})
	return v
}
