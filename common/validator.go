package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var trans ut.Translator

// ValidateRequest 校验请求参数  ShouldBindJSON方式
func ValidateRequest(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return translateValidationErrors(err)
	}
	return nil
}

// translateValidationErrors 翻译校验错误
func translateValidationErrors(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errMsgs []string
		for _, e := range validationErrors {
			errMsgs = append(errMsgs, e.Translate(trans))
		}
		return NewValidationError(strings.Join(errMsgs, "; "))
	}
	return err
}

// InitValidator 初始化校验器 自定义验证规则编写处
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册中文翻译
		zh := zh.New()
		uni := ut.New(zh, zh)
		trans, _ = uni.GetTranslator("zh")
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)

		// 注册自定义校验规则
		registerCustomValidations(v)

		// 注册获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// registerCustomValidations 注册自定义校验规则
func registerCustomValidations(v *validator.Validate) {
	// 示例1：注册手机号校验规则
	_ = v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		// 这里实现手机号校验逻辑
		mobile := fl.Field().String()
		if mobile == "" {
			return true // 允许空值（对应omitempty）
		}
		matched, _ := regexp.MatchString(`^1[3-9]\d{9}$`, mobile)
		return matched
	})

	//示例二：自定义时间格式验证函数 验证时间格式是否是2006-01-02 15:04:05 这种格式  使用示例：timeFormat
	//_ = v.RegisterValidation("timeFormat", timeFormatValidator)
}

// 自定义时间格式验证函数 验证时间格式是否是2006-01-02 15:04:05 这种格式  使用示例：timeFormat
//
//	type HeartbeatRequest struct { //请求结构体
//		HeartbeatTime string `json:"heartbeatTime" form:"heartbeatTime" binding:"omitempty,timeFormat"`
//	}
func timeFormatValidator(fl validator.FieldLevel) bool {
	timeStr := fl.Field().String()
	if timeStr == "" {
		return true // 允许空值（对应omitempty）
	}

	_, err := time.Parse("2006-01-02 15:04:05", timeStr)
	return err == nil
}
