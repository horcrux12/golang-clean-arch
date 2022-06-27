package UserService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"net/http"
)

func (service UserServiceImpl) CreateUser(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse) {
	funcName := "CreateUser"
	var inputStruct in.UserRequest
	helper.ReadFromRequestBody(request, &inputStruct)

	inputStruct.IsUpdate = false
	err := service.Validate.Struct(inputStruct)
	helper.PanicIfErrorWithLocation(err, service.FileName, funcName, ctx)

	tx, errDB := service.DB.Begin()
	helper.PanicIfErrorWithLocation(errDB, service.FileName, funcName, ctx)
	defer helper.CommitOrRollback(tx)

	userModel := service.createUserModelForInsert(inputStruct)

	// Insert data user
	userModel = service.UserRepository.CreateUser(ctx, tx, userModel)

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
	password := helper.EncryptPassword(inputStruct.Password, userSecret)
	userModel := entity.UserModel{
		Username:   sql.NullString{String: inputStruct.Username},
		Password:   sql.NullString{String: password},
		FirstName:  sql.NullString{String: inputStruct.FirstName},
		LastName:   sql.NullString{String: inputStruct.LastName},
		UserSecret: sql.NullString{String: userSecret},
	}
	return userModel
}
