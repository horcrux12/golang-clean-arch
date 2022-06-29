package UserService

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
	DB             *sql.DB
	FileName       string
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
		FileName:       "UserServiceImpl.go",
	}
}

func (service UserServiceImpl) UpdateUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse) {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) DeleteUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse) {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) UserLogin(ctx *applicationModel.ContextModel, inputRequest in.UserLoginRequest) (payload out.WebResponse) {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) GetListUser(ctx *applicationModel.ContextModel) (payload out.WebResponse) {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) DetailUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse) {
	//TODO implement me
	panic("implement me")
}

func (service UserServiceImpl) toUserResponse(userModel entity.UserModel) out.UserDetailResponse {
	return out.UserDetailResponse{
		ID:        userModel.ID.Int64,
		Username:  userModel.Username.String,
		FirstName: userModel.FirstName.String,
		LastName:  userModel.LastName.String,
	}
}
