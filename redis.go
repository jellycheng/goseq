package goseq

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jellycheng/gosupport"
)

var ctx = context.Background()
var rdbObjMap = make(map[string]*redis.Client)

func GetRedisClient(cfg RedisCfg) *redis.Client {
	k := gosupport.Md5V1(fmt.Sprintf("%s%s%s", cfg.Host, cfg.Port, cfg.Password))
	if r, ok := rdbObjMap[k]; ok {
		return r
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       gosupport.Str2Int(cfg.Db),
	})
	rdbObjMap[k] = rdb
	return rdb
}

func SetKeyValue(rdb *redis.Client, key, value string) error {
	err := rdb.Set(ctx, key, value, 0).Err()
	return err
}

func GetKeyValue(rdb *redis.Client, key string) string {
	val, _ := rdb.Get(ctx, key).Result()
	return val
}

func GetRedisClient4Json(str string) *redis.Client {
	cfg := &RedisCfg{}
	_ = gosupport.JsonUnmarshal(str, cfg)
	return GetRedisClient(*cfg)
}
