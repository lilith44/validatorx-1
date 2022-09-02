package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	_18IdCardLen      = 18
	_15LenIdCardRegex = `^[1-8]\d{7}(?:0\d|10|11|12)(?:0[1-9]|[1-2][\d]|30|31)\d{3}$`
	_18LenIdCardRegex = `^(11|12|13|14|15|21|22|23|31|32|33|34|35|36|37|41|42|43|44|45|46|50|51|52|53|54|61|62|63|64|65|71|81|82)\d{4}((((((19|20)\d{2})(0[13-9]|1[012])(0[1-9]|[12]\d|30))|(((19|20)\d{2})(0[13578]|1[02])31)|((19|20)\d{2})02(0[1-9]|1\d|2[0-8])|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))0229))\d{3})|((((\d{2})(0[13-9]|1[012])(0[1-9]|[12]\d|30))|((\d{2})(0[13578]|1[02])31)|((\d{2})02(0[1-9]|1\d|2[0-8]))|(([13579][26]|[2468][048]|0[048])0229))\d{2}))(\d|X|x)$`
)

// IdCard 18位身份证号码校验. 校验了年月日合法性，以及最后一位是否合法
func IdCard(str string) bool {
	return regexp.MustCompile(_18LenIdCardRegex).MatchString(str) && checkLastNum(str)
}

// IdCard15Len 15位身份证号码校验，校验了年月日合法性.
func IdCard15Len(str string) bool {
	return regexp.MustCompile(_15LenIdCardRegex).MatchString(str)
}

// checkIdCard 判断身份证号码是否有效(仅15位)
func checkIdCard15Len(fl validator.FieldLevel) bool {
	valid := IdCard15Len(fl.Field().String())
	return valid
}

// checkIdCard 判断身份证号码是否有效(仅18位)
func checkIdCard(fl validator.FieldLevel) bool {
	valid := IdCard(fl.Field().String())
	return valid
}

// checkLastNum 判断18位身份证号码的最后一位是否合法
func checkLastNum(idCard string) bool {
	// 根据 2^(index-1) % 11 得出位置:权重 哈希表 index从右至左算
	weightList := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码值换算表
	checkMap := map[int]int{0: 1, 2: 10, 3: 9, 4: 8, 5: 7, 6: 6, 7: 5, 8: 4, 9: 3, 10: 2, 1: 0}

	var (
		count int
	)

	for index, weight := range weightList {
		count += int(idCard[index]-'0') * weight
	}

	if int(idCard[_18IdCardLen-1]-'0') != checkMap[count%11] {
		return false
	}

	return true
}
