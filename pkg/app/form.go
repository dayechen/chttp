package app

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// BindAndValid 绑定和验证
func BindAndValid(c *gin.Context, v interface{}) error {
	var err error
	c.ShouldBind(v)
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	// 将key替换成lable里的字段
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("lable")
	})
	// 注册翻译器
	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	// 注册自定义验证方法
	_ = validate.RegisterValidation("phone", phone)
	// 将自定义验证消息翻译成中文
	validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "手机号码格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field(), fe.Field())
		return t
	})

	err = validate.Struct(v)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}

// 判断是否为手机
func phone(v validator.FieldLevel) bool {
	str := v.Field().String()
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}
