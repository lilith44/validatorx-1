package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CheckName 只允许阳间名字 不允许含有^%&',;=?$
func CheckName(name string) bool {
	regular := `^[^\^%&',;=?$\x22]+$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(name)
}

func checkName(fl validator.FieldLevel) bool {
	return CheckName(fl.Field().String())
}
