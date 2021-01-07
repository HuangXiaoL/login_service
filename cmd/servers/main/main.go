package main

import (
	"flag"
	"fmt"
	"net/http"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/http_web/user_web"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/redis/user_redis"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/mysql/user_model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var (
	configFile string
	logLevel   string
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
	//logrus.Printf("加载配置文件完成---%+v", conf)
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
	logrus.Println("日志初始化完成，日志等级为：", logLevel)
}

//initLink 初始化连接
func initLink(c *Config) {
	if err := user_model.InitMysql(c.Mysql); err != nil { //初始化数据库操作
		logrus.Panicln(err)
	}
	logrus.Println("数据库初始化完成")
	if err := user_redis.InitRedis(c.Redis); err != nil { //初始化数据库操作
		logrus.Panicln(err)
	}
	logrus.Println("缓存初始化完成")
	logrus.Println("Web服务初始化.....")
	if err := http.ListenAndServe(c.HTTP.Address, user_web.NewRouter()); err != nil {
		logrus.WithError(err).Panic("Web服务初始化.....失败")
	}
}

func main() {
	fmt.Println("main")
}
