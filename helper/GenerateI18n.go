package helper

import (
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GenerateI18NBundleTranslation(bundle *i18n.Bundle, messageID string, language string, param map[string]interface{}) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = messageID
		}
	}()

	localize := i18n.NewLocalizer(bundle, language)

	if param == nil {
		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
		})
	} else {
		output = localize.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: param,
		})
	}
	return
}

func SwitchLanguageI18N(model applicationModel.ContextModel) (language string) {
	switch model.Locale {
	case constanta.IDLangConstanta:
		language = constanta.IDLangI18NCOnstanta
		break
	case constanta.ENLangConstanta:
		language = constanta.ENLangConstanta
		break
	default:
		language = constanta.IDLangI18NCOnstanta
		break
	}

	return
}
