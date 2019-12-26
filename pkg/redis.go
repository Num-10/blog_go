package pkg

import (
	"blog_go/conf"
	"blog_go/util/e"
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"sync"
	"time"
)

var Redis *redis.Client

func RedisSetUp() {
	client := redis.NewClient(&redis.Options{
		Addr: conf.RedisIni.Addr,
		Password: conf.RedisIni.Password,
		DB: conf.RedisIni.Db,
		PoolSize: conf.RedisIni.PoolSize,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("connect redis fail: " + err.Error())
		os.Exit(e.SERVICE_CONNECT_MODEL)
	}

	Redis = client
}

// 连接池测试
func connectPoolTest() {
	fmt.Println("-----------------------welcome to connect Pool Test-----------------------")
	start := time.Now().UnixNano()
	client := Redis
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, IdleConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().IdleConns);
		}()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println(start, end)
}