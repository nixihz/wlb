package clients

import (
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

// 定义多个client的别名 用于依赖注入

type (
	MysqlClient    xorm.EngineInterface
	MysqlLogClient xorm.EngineInterface
)

type (
	StarRocksDimClient      xorm.EngineInterface
	StarRocksDmClient       xorm.EngineInterface
	StarRocksRealtimeClient xorm.EngineInterface
)

type (
	RedisClient redis.Cmdable
)
