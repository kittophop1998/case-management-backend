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

	var addr string
	var pass string
	if appcore_config.Config.Mode == "development" {
		addr = appcore_config.Config.RedisRailwayURL
		pass = appcore_config.Config.RedisRailwayPassword
	} else {
		addr = appcore_config.Config.RedisUrl
		pass = appcore_config.Config.RedisPass
	}

	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	var ctx = context.Background()

	status := Cache.Ping(ctx)
	if status.Err() != nil {
		panic("cannot connect redis database >> " + status.Err().Error())
	}
}
