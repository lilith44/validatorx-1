package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	_15IdCardLen      = 15
	_18IdCardLen      = 18
	_15LenIdCardRegex = `^[1-9]\d{7}(?:0\d|10|11|12)(?:0[1-9]|[1-2][\d]|30|31)\d{3}$`
	_18LenIdCardRegex = `^[1-9]\d{5}((((((19|20)\d{2})(0[13-9]|1[012])(0[1-9]|[12]\d|30))|(((19|20)\d{2})(0[13578]|1[02])31)|
					((19|20)\d{2})02(0[1-9]|1\d|2[0-8])|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))0229))\d{3})|
					((((\d{2})(0[13-9]|1[012])(0[1-9]|[12]\d|30))|((\d{2})(0[13578]|1[02])31)
					|((\d{2})02(0[1-9]|1\d|2[0-8]))|(([13579][26]|[2468][048]|0[048])0229))\d{2}))(\d|X|x)$`
)

// IdCard 身份证号码校验. 满足15位或18位身份证校验。18位身份证使用GB11643-1999标准做校验位验证。
func IdCard(str string) bool {
	var (
		match bool
		err   error
	)

	switch len(str) {
	case _15IdCardLen:
		if match, err = regexp.Match(_15LenIdCardRegex, []byte(str)); nil != err {
			panic(err)
		}
		return match
	case _18IdCardLen:
		if match, err = regexp.Match(_18LenIdCardRegex, []byte(str)); nil != err {
			panic(err)
		}
		if !match {
			return false
		}

		if match = checkLastNum(str); !match {
			return false
		}

	default:
		return false
	}

	return true
}

func checkLastNum(idCard string) bool {
	// 根据 2^(index-1) % 11 得出位置:权重 哈希表 index从右至左算
	weightList := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1}
	// 校验码值换算表
	checkMap := map[int]int{0: 1, 2: 10, 3: 9, 4: 8, 5: 7, 6: 6, 7: 5, 8: 4, 9: 3, 10: 2, 1: 0}

	var (
		count int
		data  int
	)
	for index, value := range idCard {
		data = int(value - '0')

		if index != _18IdCardLen-1 {
			count += data * weightList[index]
		} else {
			if value == 'x' || value == 'X' {
				data = 10
			}
		}
	}

	if data != checkMap[count%11] {
		return false
	}

	return true
}

// checkIdCard 判断身份证号码是否有效d
func checkIdCard(fl validator.FieldLevel) bool {
	valid := IdCard(fl.Field().String())
	return valid
}
