package main

import (
	"cweb/global"
	router "cweb/http/route"
	"cweb/pkg/cache"
	"cweb/pkg/file"
	"cweb/pkg/logger"
	"cweb/pkg/nosql"
	"cweb/pkg/setting"
	"cweb/pkg/sql"
	"cweb/pkg/sql/migrate"
	"fmt"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	if err := setupSetting(); err != nil {
		fmt.Printf("初始化失败 %v \n", err.Error())
	}
	if err := setupDBEngine(); err != nil {
		fmt.Printf("初始化数据库失败 %v \n", err.Error())
	}
	if err := setupLogger(); err != nil {
		fmt.Printf("初始化日志失败 %v \n", err.Error())
	}
	if err := setupRDBEngine(); err != nil {
		fmt.Printf("初始化Redis失败 %v \n", err.Error())
	}
}

func main() {
	router := router.NewRouter()
	s := &http.Server{
		Addr:         ":" + global.ServerSetting.HTTPPost,
		Handler:      router,
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	}
	s.ListenAndServe()
}

// 初始化项目配置
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// 项目配置
	if err = setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	// 项目的应用信息
	if err = setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	// JWT的信息
	if err = setting.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	// 数据库配置
	if err = setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	// websocket配置
	if err = setting.ReadSection("Socket", &global.SocketSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("Redis", &global.RedisSetting); err != nil {
		return err
	}
	// 设置缓存信息
	global.Cache = cache.NewCache()
	return nil
}

// 初始化数据库
func setupDBEngine() error {
	var err error
	global.DB, err = sql.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	// 执行迁移文件
	migrate.Run(global.DB)
	return nil
}

// 初始化日志
func setupLogger() error {
	var err error
	currentPatch, _ := file.GetCurrentPath()
	path := currentPatch + global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Log, err = logger.NewLogger(path, global.AppSetting.LogLevel, global.ServerSetting.RunMode)
	if err != nil {
		return err
	}
	global.Log.Infof("同步日志")
	return nil
}

// 初始化redis
func setupRDBEngine() error {
	var err error
	global.RDB, err = nosql.NewRDBEngine()
	if err != nil {
		return err
	}
	return nil
}
