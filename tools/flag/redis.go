package flag

import "github.com/wpf1118/toolbox/tools/env"

//RedisOpts the Mongo options.
type RedisOpts struct {
	Endpoint string
}

//NewDefaultRedisOpts returns a new default mongodb options.
func NewDefaultRedisOpts() *RedisOpts {
	return &RedisOpts{
		Endpoint: env.GetEnv(env.RedisEndpoint, "127.0.0.1:16379"),
	}
}
