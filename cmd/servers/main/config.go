package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

var cfg Config

// Config 系统配置数据类型
type Config struct {
	HTTP struct {
		Address string `toml:"address"`
	} `toml:"http"`
	Mysql struct {
		UserName string `toml:"user_name"`
		Password string `toml:"password"`
		IP       string `toml:"ip"`
		Port     string `toml:"port"`
		DbName   string `toml:"db_name"`
		MaxConn  int    `toml:"max_conn"`
	} `toml:"mysql"`
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

// Get 获得系统配置
func Get() Config {
	return cfg
}
