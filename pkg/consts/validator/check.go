package validator

import (
	validatorPkg "github.com/go-playground/validator/v10"
	"regexp"
	"unicode"
)

func checkNamePattern() validatorPkg.Func {
	return func(fl validatorPkg.FieldLevel) bool {
		value := fl.Field().String()
		result, err := regexp.MatchString("^[a-zA-Z\u4e00-\u9fa5]{1}[a-zA-Z0-9_\u4e00-\u9fa5]{0,30}$", value)
		if err != nil {
			return false
		}
		return result
	}
}

func checkIpPattern() validatorPkg.Func {
	return func(fl validatorPkg.FieldLevel) bool {
		value := fl.Field().String()
		result, err := regexp.MatchString(`^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$`, value)
		if err != nil {
			return false
		}
		return result
	}
}

func checkPasswordPattern() validatorPkg.Func {
	return func(fl validatorPkg.FieldLevel) bool {
		value := fl.Field().String()
		if len(value) < 8 || len(value) > 30 {
			return false
		}

		hasNum := false
		hasLetter := false
		for _, r := range value {
			if unicode.IsLetter(r) && !hasLetter {
				hasLetter = true
			}
			if unicode.IsNumber(r) && !hasNum {
				hasNum = true
			}
			if hasLetter && hasNum {
				return true
			}
		}

		return false
	}
}

func checkMinPattern() validatorPkg.Func {
	return func(fl validatorPkg.FieldLevel) bool {
		return true
	}
}
