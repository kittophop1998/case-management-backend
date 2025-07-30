package appcore_cache

import (
	"case-management/appcore/appcore_config"
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func InitCache(logger *slog.Logger) {
	logger.Info("Init Cache")
	Cache = redis.NewClient(&redis.Options{
		Addr:     appcore_config.Config.RedisUrl,
		Password: appcore_config.Config.RedisPass,
		DB:       0,
	})

	var ctx = context.Background()

	status := Cache.Ping(ctx)
	if status.Err() != nil {
		panic("cannot connect redis database >> " + status.Err().Error())
	}
}
