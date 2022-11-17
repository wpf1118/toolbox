package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/wpf1118/toolbox/tools/flag"
)

var redisClient *redis.Client

func RedisInit(kvOpts *flag.RedisOpts) {
	if redisClient != nil {
		return
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     kvOpts.Endpoint,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

type Redis struct {
	*redis.Client
}

func NewRedis() *Redis {
	if redisClient == nil {
		panic("kv 模块没有初始化")
	}

	return &Redis{
		redisClient,
	}
}

func (r *Redis) TTL(ctx context.Context, key string) (secs float64, err error) {
	var duration time.Duration
	duration, err = r.Client.TTL(ctx, key).Result()
	if err != nil {
		return
	}

	secs = duration.Seconds()

	return
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}) (string, error) {
	return r.Client.Set(ctx, key, value, 0).Result()
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	s, e := r.Client.Get(ctx, key).Result()
	if e != nil && e.Error() == "redis: nil" {
		return "", nil
	}

	if e != nil {
		return "", e
	}

	return s, nil
}

func (r *Redis) Del(ctx context.Context, key string) (int64, error) {
	return r.Client.Del(ctx, key).Result()
}

func (r *Redis) Push(ctx context.Context, key, value string) (int64, error) {
	return r.Client.RPush(ctx, key, value).Result()
}

func (r *Redis) LRange(ctx context.Context, key string) ([]string, error) {
	return r.Client.LRange(ctx, key, 0, -1).Result()
}

func (r *Redis) SAddStr(ctx context.Context, key string, members ...string) (err error) {
	if len(members) == 0 {
		return fmt.Errorf("members is nil")
	}

	var memberList []interface{}
	for _, v := range members {
		memberList = append(memberList, v)
	}

	return r.Client.SAdd(ctx, key, memberList...).Err()
}

func (r *Redis) SMembers(ctx context.Context, key string) (members []string, err error) {
	return r.Client.SMembers(ctx, key).Result()
}
