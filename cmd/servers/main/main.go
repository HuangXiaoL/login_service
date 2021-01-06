package main

import (
	"flag"
	"fmt"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/mysql/user_model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	configFile string
	logLevel   string
	db         *sqlx.DB
)

func init() {
	flag.StringVar(&configFile, "add", "", "config file")
	flag.StringVar(&logLevel, "log", "", "log level")
	flag.Parse()

	initLog()
	conf, err := loadConfigFile(configFile)
	if err != nil {
		logrus.WithError(err).Fatal("load config")
	}
	initLink(conf)
}

//initLog 日志初始化 默认为info 级别日志
func initLog() {
	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	fmt.Println(logLevel)
}

//initLink 初始化连接
func initLink(c *Config) {
	fmt.Println(c)
	err := user_model.InitMysql(c.Mysql)
	fmt.Println(err)
}

//initRedis 初始化Redis
func initRedis(c *Config) {

}
func main() {
	fmt.Println("main")
}
