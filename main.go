package main

import (
	"fmt"
	"gobase/global"
	"gobase/http/model"
	router "gobase/http/route"
	"gobase/pkg/logger"
	"gobase/pkg/setting"
	"log"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
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
}

func main() {
	router := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HTTPPost,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
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
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	// 数据库配置
	if err = setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	return nil
}

// 初始化数据库
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
