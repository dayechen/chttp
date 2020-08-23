package setting

import "github.com/spf13/viper"

// Setting 各项配置
type Setting struct {
	vp *viper.Viper
}

// NewSetting 读取配置信息
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
