package main

import (
	"fmt"
	"github.com/bsm/redis-lock"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	//连接redis服务器
	client := redis.NewClient(&redis.Options{
		Network:	"tcp",//以TCP的方式进行传输
		Addr:		"127.0.0.1:6379",//传输的ip地址
	})
	defer client.Close()

	//使用默认的设置获取一个新锁
	lock,err := lock.Obtain(client,"lock.foo",nil)
	if err != nil {
	   fmt.Println("Error: %s\n",err.Error())
		return
	}else if lock == nil{
		fmt.Println("Error: could not obtain lock")
		return
	}
	//最后要记得开锁
	defer lock.Unlock()

	//运行第一次
	fmt.Println("I have lock!")
	time.Sleep(200 * time.Millisecond)

	//更新锁
	ok,err := lock.Lock()
	if err != nil{
		fmt.Println("Error: %s\n",err.Error())
		return
	}else if !ok {
		fmt.Println("Error: could not renew lock")
		return
	}
	fmt.Println("I have renewed my lock!")
}
