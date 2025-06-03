package dbredis

import (
	"SimpleHTMLPage/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var tokenRedisClient *redis.Client

func GetTokenStorage() *redis.Client {
	return tokenRedisClient
}

func InitTokenStorage() error {
	tempRedisClient := redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().GetRedisAddress(),
		DB:       config.GetConfig().Redis.DB,
		Password: config.GetConfig().GetRedisPassword(),
	})

	_, err := tempRedisClient.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println(err)
		return err
	}

	tokenRedisClient = tempRedisClient
	return nil
}

func GetDataFromToken(tokenStr string) (any, error) {
	data, err := tokenRedisClient.Get(context.Background(), tokenStr).Result()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func StoreToken(tokenStr string, data any) error {
	err := tokenRedisClient.Set(context.Background(), tokenStr, data, time.Second*time.Duration(config.GetConfig().GetJWTExpireDuration())).Err()

	if err != nil {
		return err
	}

	return nil
}

func DeleteToken(tokenStr string) error {
	err := tokenRedisClient.Del(context.Background(), tokenStr).Err()

	if err != nil {
		return err
	}

	return nil
}
