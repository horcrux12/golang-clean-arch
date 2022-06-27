package helper

import (
	"fmt"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorWithLocation(err error, fileName, funcName string, contextModel *applicationModel.ContextModel) {
	if err != nil {
		CreateErrorLocation(fileName, funcName, contextModel)
		panic(err)
	}
}

//func PanicIfErrorI18N(err errorModel.ErrorExceptionModel, fileName, funcName string, contextModel *applicationModel.ContextModel) {
//	if err.Error != nil {
//		CreateErrorLocation(fileName, funcName, contextModel)
//		panic(err)
//	}
//}

func CreateErrorLocation(fileName, funcName string, contextModel *applicationModel.ContextModel) {
	contextModel.LoggerModel.Location = fmt.Sprintf(`[%s, %s]`, fileName, funcName)
}
