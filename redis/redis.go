package redis

import (
	"log"
	"sync"

	"github.com/ragul28/geo-tracking-redis/config"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	*redis.Client
}

const key = "drivers"

var once sync.Once
var redisClient *RedisClient

func GetRedisClient() *RedisClient {

	env := config.GetEnv()
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     env.Redis.RdHost,
			Password: env.Redis.RdPassword,
			DB:       0,
		})

		redisClient = &RedisClient{client}
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatalf("Redis not Connected: %v", err)
	}

	return redisClient
}

// AddDriver Loc
func (c *RedisClient) AddDriverLoc(long, lat float64, id string) {
	c.GeoAdd(
		key,
		&redis.GeoLocation{Longitude: long, Latitude: lat, Name: id},
	)
}

// Remove Driver Loc
func (c *RedisClient) RemoveDriverLoc(id string) {
	c.ZRem(key, id)
}

// Search Drivers
func (c *RedisClient) SearchDrivers(limit int, lat, long, r float64) []redis.GeoLocation {

	res, _ := c.GeoRadius(key, long, lat, &redis.GeoRadiusQuery{
		Radius:      r,
		Unit:        "km",
		WithGeoHash: true,
		WithCoord:   true,
		WithDist:    true,
		Count:       limit,
		Sort:        "ASC",
	}).Result()

	return res
}
