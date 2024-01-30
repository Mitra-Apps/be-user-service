package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func Connection() *redis.Client {
	redisConnect := os.Getenv("REDIS_SERVER")
	// Initialize Redis connection
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConnect, // Your Redis server address
		Password: "",           // No password
		DB:       0,            // Default DB
	})
	return redisClient
}
