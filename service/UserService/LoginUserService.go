package UserService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
)

func (service UserServiceImpl) UserLogin(ctx *applicationModel.ContextModel, inputRequest in.UserLoginRequest) (payload out.WebResponse) {
	funcName := "UserLogin"

	userOnDB := service.UserRepository.LoginUser(ctx, entity.UserModel{
		Username: sql.NullString{String: inputRequest.Username},
	})
	if !helper.CheckIsPasswordMatch(inputRequest.Password, userOnDB.Password.String, userOnDB.UserSecret.String) {
		helper.PanicIfError(errorModel.GenerateForbiddenAccessError(service.FileName, funcName))
	}

	// todo get token

	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: GenerateMessageI18n("LOGIN_SUCCESS", helper.SwitchLanguageI18N(*ctx)),
	}

	output := service.ToLoginResponse("Login Success")
	payload.Payload.Data = output
	return
}

func (service UserServiceImpl) ToLoginResponse(token string) out.UserLoginResponse {
	return out.UserLoginResponse{AuthorizationToken: token}
}
