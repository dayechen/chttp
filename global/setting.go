package global

import (
	"cweb/pkg/logger"
	"cweb/pkg/setting"
)

var (
	// ServerSetting 项目的启动配置
	ServerSetting *setting.ServerSetting
	// AppSetting 项目的信息配置
	AppSetting *setting.AppSetting
	// JWTSetting jwt的配置
	JWTSetting *setting.JWTSetting
	// DatabaseSetting 数据库配置文件
	DatabaseSetting *setting.DatabaseSetting
	// Logger 日志文件
	Logger *logger.Logger
)
