package UserService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

func (service UserServiceImpl) CreateUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse) {
	funcName := "CreateUser"

	err := service.Validate.Struct(inputRequest)
	helper.PanicIfErrorWithLocation(err, service.FileName, funcName, ctx)

	err = app.OpenTxConnection(ctx, app.ApplicationAttribute.DBConnection)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx.ConnectionModel.Tx)

	userModel := service.createUserModelForInsert(inputRequest)

	// Insert data user
	userModel = service.UserRepository.CreateUser(ctx, userModel)

	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: "OK",
	}

	output := service.toUserResponse(userModel)
	payload.Payload.Data = output
	return
}

func (service UserServiceImpl) createUserModelForInsert(inputStruct in.UserRequest) entity.UserModel {
	userSecret := helper.GetUUID()
	if inputStruct.Locale == "" {
		inputStruct.Locale = constanta.IDLangConstanta
	}
	password := helper.EncryptPassword(inputStruct.Password, userSecret)
	userModel := entity.UserModel{
		Username:   sql.NullString{String: inputStruct.Username},
		Password:   sql.NullString{String: password},
		FirstName:  sql.NullString{String: inputStruct.FirstName},
		LastName:   sql.NullString{String: inputStruct.LastName},
		UserSecret: sql.NullString{String: userSecret},
		Locale:     sql.NullString{String: inputStruct.Locale},
	}
	return userModel
}
