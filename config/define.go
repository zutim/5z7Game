package config

import (
	"github.com/spf13/viper"
	"go.uber.org/dig"
)

const(
	httpPortKey = "server.http-port"
	logPathKey = "server.log-path"
	jwtSignKey = "server.jwt-sign"

	dbDriverKey = "db.driver"
	dbNameKey = "db.name"
	dbHostKey = "db.host"
	dbPortKey = "db.port"
	dbPassKey = "db.pass"
	dbUserKey = "db.user"
	dbMaxIdleConnectionsKey = "db.max-idle-connections"
	dbMaxOpenConnectionsKey = "db.max-open-connections"
	dbMaxLifeTimeKey = "db.max-lifetime"

	redisHostKey = "redis.host"
	redisPortKey = "redis.port"
	redisPassKey = "redis.pass"
	redisPoolSizeKey = "redis.pool-size"
	redisMaxRetriesKey = "redis.max-retries"
	redisIdleTimeoutKey = "redis.idle-timeout"
)


var container = dig.New()

func init()  {
	viper.SetDefault(httpPortKey, 8080)
	viper.SetDefault(logPathKey, "/tmp")

	// database 默认配置
	viper.SetDefault(dbNameKey, "wuziqi")
	viper.SetDefault(dbDriverKey, "mysql")
	viper.SetDefault(dbHostKey, "127.0.0.1")
	viper.SetDefault(dbPortKey, 3306)
	viper.SetDefault(dbUserKey,"root")
	// 设置连接池默认能同时打开40个连接
	viper.SetDefault(dbMaxOpenConnectionsKey, 40)
	// 设置连接池默认能同时保存最大10个空余连接
	viper.SetDefault(dbMaxIdleConnectionsKey, 10)

	// redis 默认配置
	viper.SetDefault(redisHostKey, "127.0.0.1")
	viper.SetDefault(redisPortKey, 6379)
	viper.SetDefault(redisPoolSizeKey, 100)
	viper.SetDefault(redisMaxRetriesKey, 3)
	viper.SetDefault(redisIdleTimeoutKey, 10)
}

// ReadFromFile read from file
func ReadFromFile(path string) error {
	viper.SetConfigFile(path)

	return viper.ReadInConfig()
}
