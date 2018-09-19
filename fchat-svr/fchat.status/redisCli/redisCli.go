package redisCli

import (
	"github.com/gitfong/fLog"

	"encoding/json"
	"fmt"
	"io/ioutil"

	redigo "github.com/garyburd/redigo/redis"
)

var pool *redigo.Pool
var flog *fLog.FLogger

func init() {
	flog = fLog.New()
	if flog == nil {
		panic("in rediscli, flog.New err.")
	}

	data, err := ioutil.ReadFile("redigoCfg.json")
	if err != nil {
		flog.Fatal("readfile redigoCfg.json err:%v", err)
	}
	cfg := &redigoCfg{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		flog.Fatal("unmarshal cfg:%v,err:%v", cfg, err)
	}

	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
		if err != nil {
			return nil, err
		}
		return c, nil
	}, cfg.PoolSize)
}

func Get() redigo.Conn {
	return pool.Get()
}

type redigoCfg struct {
	RedisHost string `json:"redisHost"`
	RedisPort int    `json:"redisPort"`
	PoolSize  int    `json:"poolSize"`
	Password  string `json:"password"`
}
