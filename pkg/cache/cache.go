/*
Package cache: Provides methods for working with Redis.
Package Functionality: Connects to Redis, performs ping operations, sets and retrieves data, and invalidates the cache.
File cache.go: Contains functions for creating a connection with Redis.

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package cache

import (
	"base/pkg/config"
	"strings"
)

// Initialize Function in Cache Package
func init() {

	// If Dont Enable Cache
	// We Dont Connect With Cache
	if strings.ToLower(config.Config.GetString("ENABLE_CACHE_API")) == "false" {
		return
	}

	// Remote Cache Configuration Value
	switch strings.ToLower(config.Config.GetString("REMOTE_CACHE_DRIVER")) {
	case "redis":
		config.Config.SetDefault("REMOTE_CACHE_PORT", "6379")
		redisCfg.Host = config.Config.GetString("REMOTE_CACHE_HOST")
		redisCfg.Port = config.Config.GetString("REMOTE_CACHE_PORT")
		redisCfg.Password = config.Config.GetString("REMOTE_CACHE_PASSWORD")
		redisCfg.Name = config.Config.GetInt("REMOTE_CACHE_NAME")

		if len(redisCfg.Host) != 0 && len(redisCfg.Port) != 0 {

			// Do Redis Cache Connection
			Redis = redisConnect()
		}
	}
}
