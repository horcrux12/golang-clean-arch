package errorModel

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"regexp"
)

func AddTranslation(validate *validator.Validate, validateModel []ErrorValidatorModel) {
	for i := 0; i < len(validateModel); i++ {
		validate.RegisterValidation(validateModel[i].Tag, validateModel[i].ValidateFunc)
		addTranslator(validateModel[i].Tag, validateModel[i].Messages, validate)
	}
}

func addTranslator(tag string, messages map[string]string, validate *validator.Validate) {
	for key, value := range messages {

		registerFn := func(ut ut.Translator) error {
			return ut.Add(tag, value, false)
		}

		transFn := func(ut ut.Translator, fe validator.FieldError) string {
			param := fe.Param()
			tag := fe.Tag()

			t, err := ut.T(tag, fe.Field(), param)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		}

		trans := GetTranslatorForTranslate(key)

		_ = validate.RegisterTranslation(tag, trans, registerFn, transFn)
	}
}

func GetTranslatorForTranslate(locale string) (trans ut.Translator) {
	switch locale {
	case constanta.IDLangConstanta:
		trans = app.ApplicationAttribute.IDTranslator
		break
	case constanta.ENLangConstanta:
		trans = app.ApplicationAttribute.ENTranslator
		break
	}
	return
}

func validateUsername(fl validator.FieldLevel) bool {
	usernameRegex := regexp.MustCompile("^[a-z][a-z0-9_.]+$")
	return usernameRegex.MatchString(fl.Field().String())
}
