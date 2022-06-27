package service

import (
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
)

type AbstractService struct {
	FileName    string
	ServiceName string
}

func (service AbstractService) GetCommonResponseMessage(messageID string, ctx *applicationModel.ContextModel) (output out.PayloadStatusResponse) {
	param := make(map[string]interface{})
	param["SERVICE_NAME"] = helper.GenerateI18NBundleTranslation(app.ApplicationAttribute.ConstantaBundleI18N, service.ServiceName, helper.SwitchLanguageI18N(*ctx), nil)

	output.Code = 200
	output.Message = helper.GenerateI18NBundleTranslation(app.ApplicationAttribute.CommonMessagesBundleI18N, messageID, helper.SwitchLanguageI18N(*ctx), param)
	return
}
