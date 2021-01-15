package tools

import (
	"github.com/garyburd/redigo/redis"
	"log"
)
var RdEnige redis.Conn
func InitRedisEngine(rdConfig RedisConfig)  {
	conn := rdConfig.Host+":"+rdConfig.Port
	engine,err:=redis.Dial("tcp",conn)
	if err !=nil{
		log.Fatalln("conn redis fail")
	}
	RdEnige = engine
}
