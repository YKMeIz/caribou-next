package cache

import (
	"github.com/YKMeIz/logc"
	"github.com/go-redis/redis"
	"os"
	"time"
)

var client *redis.Client

const StandTTL = time.Duration(24 * time.Hour)

func init() {
	if os.Getenv("REDIS_ADDRESS") == "" {
		logc.Error("environment variable REDIS_ADDRESS is not defined")
	}

	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		logc.Error(err)
	}

	client.ConfigSet("maxmemory", "500mb")
	client.ConfigSet("maxmemory-policy", "volatile-lru")

	logc.Important("redis ", client.ConfigGet("maxmemory-policy").String())
	logc.Important("redis ", client.ConfigGet("maxmemory").String())
}
