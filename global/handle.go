package global

import (
	"cweb/pkg/cache"
	"cweb/pkg/socket/wslogic"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	Socket *wslogic.Engine    // socket
	DB     *gorm.DB           // 数据库
	RDB    *redis.Client      // redis
	Log    *zap.SugaredLogger // 日志
	Cache  *cache.Engine      // 缓存
)
