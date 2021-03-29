package global

import (
	"cweb/pkg/logger"
	"cweb/pkg/socket/logic"

	"github.com/jinzhu/gorm"
)

var (
	Socket *logic.Engine // socket
	DB     *gorm.DB      // 数据库
	Log    *logger.Logger
)
