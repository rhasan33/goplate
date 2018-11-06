package conn

import (
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/rhasan33/goplate/config"
)

// RedisClient holds the redis client instance
type RedisClient struct {
	*redis.Client
}

// Redis is an instance *redis.Client{}
var redisCl RedisClient

// Setup assigns redis.Client interface based on config to RedisClient
func (r *RedisClient) Setup(cfg *config.RedisConf) {
	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":6379",
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := c.Ping().Err(); err != nil {
		log.Fatalln("Failed to connect redis : ", err)
		os.Exit(-1)
	}

	log.Println("Redis connection successful")
	r.Client = c
}

// ConnectRedis provides a connector to redis based on configurations set
func ConnectRedis() error {
	cfg := config.Redis()
	redisCl.Setup(cfg)
	return nil
}

// DefaultRedis returns the default RedisClient currently in Use
func DefaultRedis() RedisClient {
	return redisCl
}

// SetToRedis sets redis value
func SetToRedis(key, value string) error {
	r := DefaultRedis()
	return r.Set(config.Redis().Prefix+key, value, config.Redis().PinResetTTL).Err()
}

// GetFromRedis gets value from redis
func GetFromRedis(key string) (string, error) {
	r := DefaultRedis()
	return r.Get(config.Redis().Prefix + key).Result()
}

// DelFromRedis deletes key from redis
func DelFromRedis(key string) error {
	r := DefaultRedis()
	return r.Del(config.Redis().Prefix + key).Err()
}
