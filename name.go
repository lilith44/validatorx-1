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

func ChineseOrEnglishOrNumbers(name string) bool {
	regular := `^[A-Za-z0-9\p{Han}]+$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(name)
}

func chineseOrEnglishOrNumbers(fl validator.FieldLevel) bool {
	return ChineseOrEnglishOrNumbers(fl.Field().String())
}
