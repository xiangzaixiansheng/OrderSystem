package redis

import (
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

var (
	RdsConn *redis.Client
	once    sync.Once
)

//sync.Once 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等。作用与 init 函数类似，但有区别。

//init 函数是当所在的 package 首次被加载时执行，若迟迟未被使用，则既浪费了内存，又延长了程序加载时间。
//sync.Once 可以在代码的任意位置初始化和调用，因此可以延迟到使用时再执行，并发场景下是线程安全的。

func init() {
	once.Do(func() {
		redisConfig, _ := beego.AppConfig.GetSection("redis")
		db_num, _ := strconv.Atoi(redisConfig["db_index"])
		RdsConn = redis.NewClient(&redis.Options{
			Addr:     redisConfig["address"],
			Password: redisConfig["password"],
			DB:       db_num,
		})
		_, err := RdsConn.Ping().Result()
		if err != nil {
			logs.Info("connect redis fail", err)
			panic(err)
		}
	})

}
