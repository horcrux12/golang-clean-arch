package UserService

import (
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/helper"
)

func GenerateMessageI18n(messagesID, lang string) (output string) {
	return helper.GenerateI18NBundleTranslation(app.ApplicationAttribute.UserBundleI18N, messagesID, lang, nil)
}
