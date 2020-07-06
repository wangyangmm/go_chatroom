package main
import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var g_pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	g_pool = &redis.Pool {
		MaxIdle : maxIdle, //最大空闲连接数
		MaxActive : maxActive, //和数据库的最大连接数， 0表示没有限制
		IdleTimeout : idleTimeout, //最大空闲时间，超过这个时间，就表示空闲了
		Dial : func() (redis.Conn, error) { //初始化连接，连接到哪个ip
			return redis.Dial("tcp", address)
		},
	}
}