package global

import (
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
	// SocketSetting socket配置
	SocketSetting *setting.SocketSetting
	// Logger 日志配置
)
