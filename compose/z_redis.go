package compose

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/sohaha/gconf"

	"github.com/sohaha/zlsgo/zutil"
)

type (
	stRedisConf struct {
		Host     string
		Port     int
		Password string
		DBNumber int `mapstructure:"db"`
	}
)

func (*stRedisConf) ConfName(key ...string) string {
	if len(key) > 0 {
		return "redis." + key[0]
	}
	return "redis"
}

// noinspection GoUnusedGlobalVariable
var (
	// Redis Redis 实例
	Redis     *redis.Client
	redisConf stRedisConf
)

func (*stCompose) RedisDefaultConf(cfg *gconf.Confhub) {
	for k, v := range map[string]interface{}{
		"host":     "localhost",
		"port":     "6379",
		"password": "",
		// "db":       1,
	} {
		cfg.SetDefault(redisConf.ConfName()+"."+k, v)
	}
}

func (*stCompose) RedisReadConf(cfg *gconf.Confhub) error {
	redisConf.DBNumber = 0
	return cfg.Core.UnmarshalKey(redisConf.ConfName(), &redisConf)
}

func (*stCompose) RedisDone() {
	c, err := conn(redisConf)
	zutil.CheckErr(err)
	Redis = c
}

func conn(RedisConfig stRedisConf) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port),
		Password:    RedisConfig.Password,
		DB:          RedisConfig.DBNumber,
		IdleTimeout: time.Second * 60,
		MaxRetries:  2,
	})

	_, err := cli.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect redis, got error %w", err)
	}

	return cli, nil
}
