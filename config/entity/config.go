/*
 * @Description:配置文件对应实体
 * @Author: chiwenda
 * @Date: 2022-10-15 20:43:05
 * @LastEditTime: 2022-10-23 17:05:09
 * @LastEditors: chiwenda
 * @FilePath: /go-sombrero/config/entity/config.go
 */
package entity

//配置
type Config struct {
	Servers
	Mysql
}

//服务器配置内容
type Servers struct {
	Name string `toml:"name"`
	Port int    `toml:"port"` //端口号
}
type Mysql struct {
	Url string `toml:"url"` //连接字符串

}
