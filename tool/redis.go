package tool

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient    *redis.Pool
	REDIS_HOST     string
	REDIS_DB       int
	REDIS_PASSWORD string
)

func Redisinit() *redis.Pool {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	REDIS_PASSWORD = beego.AppConfig.String("redis.password")

	fmt.Println(REDIS_HOST)
	fmt.Println(REDIS_DB)
	fmt.Println(REDIS_PASSWORD)

	// 建立连接池
	return &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 100),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 1000),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)

			if err != nil {
				return nil, err
			}
			if REDIS_PASSWORD != "" {
				if _, err := c.Do("AUTH", REDIS_PASSWORD); err != nil {
					c.Close()
					return nil, err
				}
			}

			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
