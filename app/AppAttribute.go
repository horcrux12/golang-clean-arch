package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/horcrux12/clean-rest-api-template/config"
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"os"
	"strconv"
	"time"
)

var ApplicationAttribute applicationAttribute

type applicationAttribute struct {
	DBConnection             *sql.DB
	Validate                 *validator.Validate
	Uni                      *ut.UniversalTranslator
	IDTranslator             ut.Translator
	ENTranslator             ut.Translator
	ErrorBundleI18N          *i18n.Bundle
	CommonMessagesBundleI18N *i18n.Bundle
	ConstantaBundleI18N      *i18n.Bundle
	UserBundleI18N           *i18n.Bundle
	CronScheduler            *gocron.Scheduler
}

func GenerateApplicationAttribute() {
	dbAddress := config.ApplicationConfiguration.GetPostgreSQLAddress()
	dbSchema := config.ApplicationConfiguration.GetPostgreSQLDefaultSchema()
	dbMaxIdle := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnection()
	dbMaxOpen := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnection()
	ApplicationAttribute.DBConnection = NewDB(dbAddress, dbSchema, dbMaxOpen, dbMaxIdle)

	// Setting validation
	ApplicationAttribute.Validate = validator.New()

	ApplicationAttribute.ENTranslator = getTranslator(constanta.ENLangConstanta)
	ApplicationAttribute.IDTranslator = getTranslator(constanta.IDLangConstanta)

	t := time.Now()
	zone, offset := t.Zone()

	fmt.Println(zone, offset)
	ApplicationAttribute.CronScheduler = gocron.NewScheduler(time.FixedZone(zone, offset))
	ApplicationAttribute.CronScheduler.StartAsync()

	loadBundleI18n()
}

func getTranslator(locale string) (trans ut.Translator) {
	switch locale {
	case constanta.IDLangConstanta:
		trans = idTranslator()
		break
	case constanta.ENLangConstanta:
		trans = enTranslator()
		break
	}
	return
}

func enTranslator() ut.Translator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator(constanta.ENLangConstanta)

	en_translations.RegisterDefaultTranslations(ApplicationAttribute.Validate, trans)
	return trans
}

func idTranslator() ut.Translator {
	idLang := id.New()
	uni := ut.New(idLang, idLang)

	trans, _ := uni.GetTranslator(constanta.IDLangConstanta)
	id_translations.RegisterDefaultTranslations(ApplicationAttribute.Validate, trans)
	return trans
}

var fileNumber = 0

func loadBundleI18n() {
	prefixPath := "./messages"

	//------------ error bundle
	ApplicationAttribute.ErrorBundleI18N = i18n.NewBundle(language.Indonesian)
	ApplicationAttribute.ErrorBundleI18N.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err := ApplicationAttribute.ErrorBundleI18N.LoadMessageFile(prefixPath + "/error_messages/en-US.json")
	fileNumber++
	readError(err)

	_, err = ApplicationAttribute.ErrorBundleI18N.LoadMessageFile(prefixPath + "/error_messages/id-ID.json")
	fileNumber++
	readError(err)

	//------------ common bundle
	ApplicationAttribute.CommonMessagesBundleI18N = i18n.NewBundle(language.Indonesian)
	ApplicationAttribute.CommonMessagesBundleI18N.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = ApplicationAttribute.CommonMessagesBundleI18N.LoadMessageFile(prefixPath + "/common_messages/en-US.json")
	fileNumber++
	readError(err)

	_, err = ApplicationAttribute.CommonMessagesBundleI18N.LoadMessageFile(prefixPath + "/common_messages/id-ID.json")
	fileNumber++
	readError(err)

	//------------ constanta bundle
	ApplicationAttribute.ConstantaBundleI18N = i18n.NewBundle(language.Indonesian)
	ApplicationAttribute.ConstantaBundleI18N.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = ApplicationAttribute.ConstantaBundleI18N.LoadMessageFile(prefixPath + "/common_messages/constanta/en-US.json")
	fileNumber++
	readError(err)

	_, err = ApplicationAttribute.ConstantaBundleI18N.LoadMessageFile(prefixPath + "/common_messages/constanta/id-ID.json")
	fileNumber++

	//------------ user bundle
	ApplicationAttribute.UserBundleI18N = i18n.NewBundle(language.Indonesian)
	ApplicationAttribute.UserBundleI18N.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, err = ApplicationAttribute.UserBundleI18N.LoadMessageFile(prefixPath + "/user_messages/en-US.json")
	fileNumber++
	readError(err)

	_, err = ApplicationAttribute.UserBundleI18N.LoadMessageFile(prefixPath + "/user_messages/id-ID.json")
	fileNumber++
	readError(err)
}

func readError(err error) {
	if err != nil {
		fmt.Println(err.Error() + " at file " + strconv.Itoa(fileNumber))
		os.Exit(3)
	}
}
