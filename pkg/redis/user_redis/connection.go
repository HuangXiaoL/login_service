package user_redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	IP          string        `toml:"ip"`
	Port        string        `toml:"port"`
	MaxIdle     int           `toml:"max_idle"`
	MaxActive   int           `toml:"max_active"`
	IdleTimeout time.Duration `toml:"idle_timeout"`
}

var (
	Pool redis.Pool
	conn redis.Conn
)

//initRedis 初始化Redis
func InitRedis(r Redis) (err error) {
	Pool = redis.Pool{
		MaxIdle:     r.MaxIdle,
		MaxActive:   r.MaxActive,
		IdleTimeout: r.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", r.IP+":"+r.Port)
		},
	}
	return err
}
