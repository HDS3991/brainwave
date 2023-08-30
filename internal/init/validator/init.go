package validator

import (
	"brainwave/internal/global"
	brainwaveValidator "brainwave/pkg/consts/validator"
	"errors"
	"fmt"
	validatorPkg "github.com/go-playground/validator/v10"
)

func Init() {
	validator := validatorPkg.New()
	str := ""
	for i := brainwaveValidator.Min; i < brainwaveValidator.Max; i++ {
		str = i.String()
		ruleLoader, ok := brainwaveValidator.RuleLoaderMap[str]
		if !ok {
			panic(errors.New(fmt.Sprintf("rule loader not found, %s", str)))
		}
		switch loader := ruleLoader.(type) {
		case validatorPkg.Func:
			if err := validator.RegisterValidation(str, loader); err != nil {
				panic(err)
			}
		}

	}

	global.VALID = validator
}
