package user

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	IP           string `toml:"ip"`
	Port         int    `toml:"port"`
	Password     string `toml:"password"`
	DB           int    `toml:"db"`
	PoolSize     int    `toml:"pool_size"`
	MinIdleConns int    `toml:"min_idle_conns"`
}

var (
	client *redis.Client
	Nil    = redis.Nil
)

//initRedis 初始化Redis
func InitRedis(r Redis) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", r.IP, r.Port),
		Password:     r.Password, // no password set
		DB:           r.DB,       // use default DB
		PoolSize:     r.PoolSize,
		MinIdleConns: r.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
