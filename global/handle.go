package global

import (
	"cweb/pkg/socket/logic"

	"github.com/jinzhu/gorm"
)

var (
	Socket *logic.Engine // socket 工具
	DB     *gorm.DB      // 数据库工具
)
