package goseq

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/jellycheng/gosupport"
)

type MyRedisClient struct {
	rdb *redis.Client
	cfg RedisCfg
}

func (m MyRedisClient) GetCfg() RedisCfg {
	return m.cfg
}

func (m MyRedisClient) GetRedisClient() *redis.Client {
	return m.rdb
}

var ctx = context.Background()
var rdbObjMap = make(map[string]*MyRedisClient)

func NewRedisClient(cfg RedisCfg) *MyRedisClient {
	myRedis := &MyRedisClient{
		cfg: cfg,
	}
	k := gosupport.Md5V1(fmt.Sprintf("%s%s%s%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password))
	if r, ok := rdbObjMap[k]; ok {
		return r
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       gosupport.Str2Int(cfg.Db),
	})
	myRedis.rdb = rdb
	rdbObjMap[k] = myRedis
	return myRedis
}

func SetKeyValue(myRedis *MyRedisClient, key, value string) error {
	tmpKey := fmt.Sprintf("%s%s", myRedis.cfg.Prefix, key)
	err := myRedis.rdb.Set(ctx, tmpKey, value, 0).Err()
	return err
}

func GetKeyValue(myRedis *MyRedisClient, key string) string {
	tmpKey := fmt.Sprintf("%s%s", myRedis.cfg.Prefix, key)
	val, _ := myRedis.rdb.Get(ctx, tmpKey).Result()
	return val
}

// 前缀+年月日+今天过去秒+今天自增序列号
func DefaultSeq(myRedis *MyRedisClient, seqPrefix string) string {
	ret := ""
	tmpKey := GetSeqDefaultRedisKey()
	num := myRedis.rdb.Incr(ctx, tmpKey).Val()
	incrStr := strconv.FormatInt(num%999999, 10)
	ret = fmt.Sprintf("%s%s%d%s", seqPrefix, gosupport.TimeNow2Format("20060102"), TodayPastTime(), incrStr)
	return ret
}

func NewRedisClient4Json(str string) *MyRedisClient {
	cfg := &RedisCfg{}
	_ = gosupport.JsonUnmarshal(str, cfg)
	return NewRedisClient(*cfg)
}

func GetSeqDefaultRedisKey() string {
	return fmt.Sprintf("goseq:%s", gosupport.TimeNow2Format("20060102"))
}
