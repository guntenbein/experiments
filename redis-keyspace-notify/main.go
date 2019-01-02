package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("redis connection ok\n")
	err = redisClient.Do("CONFIG", "SET", "notify-keyspace-events", "KEA").Err()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	ps := redisClient.PSubscribe("__key*__:set")
	go func() {
		for {
			m, err := ps.ReceiveMessage()
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				return
			}
			fmt.Printf("received message:\n   payload: %s\n   pattern: %s\n   channel: %s\n",
				m.Payload, m.Pattern, m.Channel)
		}
	}()
	time.Sleep(1 * time.Second)
	redisClient.Do("SET", "samplekey", "ipsum")
	redisClient.Do("SET", "samplekey", "volum")
	time.Sleep(1 * time.Second)
}
