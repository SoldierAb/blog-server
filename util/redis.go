package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

const(
	RedisUrl = "redis://127.0.0.1:6379"
	RedisPass = "redisgjchen"
)

func  NewRedisPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle:3,
		IdleTimeout:240,
		Dial: func() (redis.Conn, error) {
			c,err := redis.DialURL(RedisUrl)
			if err !=nil{
				return nil,fmt.Errorf("redis conn error : %s ",err)
			}
			//验证密码
			if _,authErr :=  c.Do("AUTH",RedisPass);authErr!=nil{
				return nil,fmt.Errorf("redis auth password error: &s ",err)
			}
			return c,nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("PING")
			if err !=nil{
				return fmt.Errorf("ping redis error : %s  ",err)
			}
			return nil
		},
	}
}


func SetRedisKey(key,value string,exType string,sec int) error{

	redisConn := NewRedisPool().Get()
	defer redisConn.Close()

	_,err := redisConn.Do("Set",key,value,exType,sec)
	if err !=nil{
		log.Fatalf("SetRedisKey error : %v",err)
		return err
	}

	return nil
}


func GetRedisValue(key string) (string,error){
	redisConn :=  NewRedisPool().Get()

	defer redisConn.Close()

	redisVal,err  := redis.String(redisConn.Do("Get",key))

	if err !=nil{
		log.Fatalf("GetRedisValue error : %v",err)
		return "",err
	}

	return redisVal,nil

}

func SetKeyExpire(k string, ex int) error{

	redisConn := NewRedisPool().Get()

	defer redisConn.Close()
	_, err := redisConn.Do("EXPIRE", k, ex)
	if err != nil {
		log.Fatalf("SetKeyExpire Error: &v" , err)
		return err
	}
	return nil
}

func DelRedisKey(k string) error {
	redisConn := NewRedisPool().Get()
	defer redisConn.Close()
	_, err := redisConn.Do("DEL", k)
	if err != nil {
		log.Fatalf("DelRedisKey Error: &v" , err)
		return err
	}
	return nil
}


func SetJson(k string, data interface{}) error {
	c := NewRedisPool().Get()
	defer c.Close()
	value, _ := json.Marshal(data)
	n, _ := c.Do("SETNX", k, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

func getJsonByte(key string) ([]byte, error) {
	c := NewRedisPool().Get()
	defer c.Close()
	jsonGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonGet, nil
}

