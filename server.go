package main

import (
  "github.com/garyburd/redigo/redis"
  "fmt"
  "time"
  "os"
)

func init() {
  linkInit()
}

var RedisPool *redis.Pool

func RedisConnection() redis.Conn {

  if RedisPool == nil {
    RedisPool = &redis.Pool{
      MaxIdle: 3,
      IdleTimeout: 240 * time.Second,
      Dial: func() (redis.Conn, error) {
        address := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
        c, err := redis.Dial("tcp", address)
        if err != nil {
          // log here
        }
        return c, nil
      },
    }
  }
  
  return RedisPool.Get()
}

func main() {

  fmt.Println("start server...")
}



