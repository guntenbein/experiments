package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"os"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pong)

	kk, err := client.Keys("d*").Result()
	fmt.Printf("kk %+v\n",kk)

	client.SAdd("1", 1,2,3)
	client.SAdd("2", 1,2)
	client.SAdd("3",1,2)

	ii, _ := client.SInter("1","2","3").Result()

	fmt.Printf("the result: %+v\n",ii)

	client.H
}
