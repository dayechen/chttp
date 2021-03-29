package model

import (
	"cweb/global"
	"cweb/pkg/setting"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Model 模型的公用参数
type Model struct {
	ID        int   `gorm:"primary_key" json:"id"`
	CreatedOn int64 `json:"created_on"`
	DeletedOn int64 `json:"deleted_on"`
}

// NewDBEngine 创建数据库实例
func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	createFied, _ := scope.FieldByName("CreatedOn")
	createFied.Set(time.Now().Unix())
}
