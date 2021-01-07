package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/redis/user_redis"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/mysql/user_model"

	"github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

var cfg Config

// Config 系统配置数据类型
type Config struct {
	HTTP struct {
		Address string `toml:"address"`
	} `toml:"http"`
	user_model.Mysql `toml:"user_mysql"` //用户数据库的连接MySQL需要的结构参数
	user_redis.Redis `toml:"user_redis"` //用户数据库的连接MySQL需要的结构参数
}

//loadConfigFile 配置文件相关信息载入进来 add 配置文件的路径
func loadConfigFile(add string) (c *Config, err error) {
	logrus.Println("开始加载配置文件，路径为---", add)
	if err = loadFile(add); err != nil {
		return
	}
	c = &cfg
	return
}

// LoadFile 从文件载入配置信息
func loadFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read file, %w", err)
	}

	return toml.Unmarshal(data, &cfg)
}
