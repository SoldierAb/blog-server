package util

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	conn redis.Conn
}

var RedisInstance *Redis

func InitConnectRedis() error{
	redisConn,err := redis.Dial("tcp","127.0.0.1:6379")

	if err !=nil{
		fmt.Println("connect to redis error")
		return err
	}

	//defer redisConn.Close()

	RedisInstance = &Redis{
		conn:redisConn,
	}

	return nil
}


func SetRedisKey(key,value string) error{
	_, err := RedisInstance.conn.Do("SET", key, value, "EX", "2")
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	return nil
}


func GetRedisKey(key string) (string,error){

	res, err := redis.String(RedisInstance.conn.Do("GET", key))

	if err != nil {
		fmt.Println("redis get failed:", err)
		return "",err
	}

	fmt.Printf("Get mykey: %v \n", res)
	return res,nil
}


func CloseRedis() error{
	return RedisInstance.conn.Close()
}

