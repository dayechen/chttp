package setting

import "time"

// ServerSetting 服务器配置
type ServerSetting struct {
	RunMode      string
	HTTPPost     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppSetting 项目的配置信息
type AppSetting struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
	LogLevel    string
}

// JWTSetting jwt 的项目配置
type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
	Salt   string
}

// DatabaseSetting 数据库设置
type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type SocketSetting struct {
	Active bool
	Url    string
}

// ReadSection 读取配置文件
func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
