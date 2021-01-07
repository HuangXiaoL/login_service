package user_model

import (
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	UserName string `toml:"user_name"`
	Password string `toml:"password"`
	IP       string `toml:"ip"`
	Port     string `toml:"port"`
	DbName   string `toml:"db_name"`
	MaxConn  int    `toml:"max_conn"`
}

var db *sqlx.DB

//initMysql 初始化MySQL
func InitMysql(c Mysql) (err error) {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//logrus.Printf("正在初始化数据库，参数为----%+v", c)
	path := strings.Join([]string{c.UserName, ":", c.Password, "@tcp(", c.IP, ":", c.Port, ")/", c.DbName, "?charset=utf8mb4"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err = sqlx.Open("mysql", path)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(c.MaxConn) // 最大连接数
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(5 * time.Minute)
	//验证连接
	if err = db.Ping(); err != nil {
		return
	}

	return nil
}
