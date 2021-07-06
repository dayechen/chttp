package global

import (
	"cweb/pkg/cache"
	"cweb/pkg/logger"
	"cweb/pkg/socket/logic"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	Socket *logic.Engine  // socket
	DB     *gorm.DB       // 数据库
	RDB    *redis.Client  // redis
	Log    *logger.Engine // 日志
	Cache  *cache.Engine  // 缓存
)
