package ginvalid

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 密码校验器
var passwordValidator validator.Func = func(fl validator.FieldLevel) bool {
	pwd, ok := fl.Field().Interface().(string)
	if ok {
		match, _ := regexp.MatchString("^[0-9a-zA-Z_-]{6,18}$", pwd)
		return match
	}
	return false // (why gin doc demo reture true)
}

func RegisterValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", passwordValidator)
	}
}
