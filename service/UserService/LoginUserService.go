package UserService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
)

func (service UserServiceImpl) UserLogin(ctx *applicationModel.ContextModel, inputRequest in.UserLoginRequest) (payload out.WebResponse) {
	// Validate login on DB
	_ = service.validateLoginOnDB(ctx, inputRequest)

	// todo get token

	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: GenerateMessageI18n("LOGIN_SUCCESS", helper.SwitchLanguageI18N(*ctx)),
	}

	output := service.ToLoginResponse("Login Success")
	payload.Payload.Data = output
	return
}

func (service UserServiceImpl) validateLoginOnDB(ctx *applicationModel.ContextModel, inputRequest in.UserLoginRequest) (userOnDB entity.UserModel) {
	funcName := "validateLoginOnDB"

	// Open Tx Connection
	err := app.OpenTxConnection(ctx, app.ApplicationAttribute.DBConnection)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx.ConnectionModel.Tx)

	// get user from DB
	userOnDB = service.UserRepository.LoginUser(ctx, entity.UserModel{
		Username: sql.NullString{String: inputRequest.Username},
	})
	if !helper.CheckIsPasswordMatch(inputRequest.Password, userOnDB.Password.String, userOnDB.UserSecret.String) {
		panic(errorModel.GenerateForbiddenAccessError(service.FileName, funcName))
	}

	return
}

func (service UserServiceImpl) ToLoginResponse(token string) out.UserLoginResponse {
	return out.UserLoginResponse{AuthorizationToken: token}
}
