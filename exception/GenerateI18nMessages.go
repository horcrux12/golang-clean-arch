package exception

import (
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateErrorI18NMessage(err errorModel.ErrorExceptionModel, language string) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = err.Error.Error()
		}
	}()

	localize := i18n.NewLocalizer(app.ApplicationAttribute.ErrorBundleI18N, language)
	if err.ErrorParameters == nil {
		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID: err.Error.Error(),
		})
	} else {
		param := make(map[string]interface{})
		for i := 0; i < len(err.ErrorParameters); i++ {
			var parameterValue = err.ErrorParameters[i].ParamValue
			if err.ErrorParameters[i].ParamKey == "FieldName" {
				parameterValue = GenerateI18nConstanta(err.ErrorParameters[i].ParamValue, language, nil)
			}
			param[err.ErrorParameters[i].ParamKey] = parameterValue
			if param["RuleName"] != nil {
				param["RuleName"] = GenerateI18nConstanta(param["RuleName"].(string), language, nil)
			}
		}
	}
	return
}

func GenerateI18nConstanta(messageID string, language string, param map[string]interface{}) string {
	return helper.GenerateI18NBundleTranslation(app.ApplicationAttribute.ConstantaBundleI18N, messageID, language, param)
}
