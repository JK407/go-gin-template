package cache

import (
	"context"
	"fmt"
	"gin-template/conf"
	redis "github.com/go-redis/redis/v8"
	"github.com/gogf/gf/util/gconv"
	logx "github.com/yunduan16/micro-service-go-component-log"
	"github.com/yunduan16/micro-service-go-component-redisx"
	"log"
	"time"
)

type RedisConnection struct {
	Cfg    *conf.RedisConfig
	client *redis.Client
	close  func()
}

type RedisCache struct {
	KeyPrefix  string
	ExpireTime time.Duration
	Conn       *RedisConnection
}

var RedisConnections map[string]*RedisConnection

func InitRedis(cfg *conf.Config) func() {
	RedisConnections = make(map[string]*RedisConnection)
	for s, conf := range cfg.Redis {
		client, cls, err := redisx.InitRedis(conf.ConnConf)
		if err != nil {
			log.Fatalf("redis config :%+v connect fail, error :%v", conf.ConnConf, err)
		}
		RedisConnections[s] = &RedisConnection{
			client: client,
			close:  cls,
			Cfg:    conf,
		}
	}
	return closeRedis()
}

func closeRedis() func() {
	return func() {
		for s, connection := range RedisConnections {
			logx.Info(logx.Fields{}, fmt.Sprintf("redis connection[%s] is closed", s))
			connection.close()
		}
	}
}

func (c *RedisCache) Get(ctx context.Context, suffixKey string) interface{} {
	cacheKey := c.KeyPrefix + suffixKey
	val, err := c.Conn.client.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Fatalf("redis get error:%v", err)
		return nil
	} else {
		return val
	}
}

func (c *RedisCache) GetByKey(ctx context.Context, key string) interface{} {
	val, err := c.Conn.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Fatalf("redis get error:%v", err)
		return nil
	} else {
		return val
	}
}

func (c *RedisCache) GetTTL(ctx context.Context, key string) int64 {
	val, err := c.Conn.client.TTL(ctx, key).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		log.Fatalf("redis getTTL error:%v", err)
		return 0
	} else {
		return gconv.Int64(val.Seconds())
	}
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}) error {
	cacheKey := c.KeyPrefix + key
	//fmt.Printf("set %s => %s\n", cacheKey, value)
	if err := c.Conn.client.Set(ctx, cacheKey, value, c.ExpireTime).Err(); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *RedisCache) Del(ctx context.Context, key string) error {
	cacheKey := c.KeyPrefix + key
	//fmt.Printf("set %s => %s\n", cacheKey, value)
	if err := c.Conn.client.Del(ctx, cacheKey).Err(); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *RedisCache) WithKeyPrefix(keyPrefix string) *RedisCache {
	c.KeyPrefix = keyPrefix
	return c
}

func (c *RedisCache) GetKeyPrefix() string {
	return c.KeyPrefix
}

func (c *RedisCache) SetPipeline(ctx context.Context, pipe redis.Pipeliner, key string, value interface{}) redis.Pipeliner {
	if pipe == nil {
		pipe = c.Conn.client.Pipeline()
	}
	cacheKey := c.KeyPrefix + key
	pipe.Set(ctx, cacheKey, value, c.ExpireTime)
	//fmt.Printf("set %s => %s\n", cacheKey, value)
	return pipe
}

func (c *RedisCache) DelPipeline(ctx context.Context, pipe redis.Pipeliner, key string) redis.Pipeliner {
	if pipe == nil {
		pipe = c.Conn.client.Pipeline()
	}
	pipe.Del(ctx, key)
	//fmt.Printf("set %s => %s\n", cacheKey, value)
	return pipe
}

func (c *RedisCache) ExecPipeline(ctx context.Context, pipe redis.Pipeliner) error {
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("redis exec cmds: %+v fail, error:%v", cmds, err)
	}
	return nil
}

//// GetRedisCache 由于yaml解析存在将map键值转为小写问题，所以这里也要处理一下
//func GetRedisCache(key string) *RedisCache {
//	key = strings.ToLower(key)
//	//fmt.Println(RedisConnections)
//	if _, ok := RedisConnections[key]; !ok {
//		log.Fatalf(fmt.Sprintf("redis[%s] config is not exist!", key))
//	}
//	switch key {
//	case "article":
//		return NewArticleCache(RedisConnections[key])
//	}
//	return nil
//}

func (c *RedisCache) Scan(ctx context.Context, cursor uint64, count int64) ([]string, uint64, error) {
	//当数据量特别大时，采用正则匹配方式查找太慢，而且不准确
	return c.Conn.client.Scan(ctx, cursor, "", count).Result()
}
