package repository

import (
	"fmt"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
)

type UserRepositoryImpl struct {
	AbstractRepository
}

func NewUserRepository() UserRepository {
	repo := UserRepositoryImpl{}
	repo.TableName = `"user"`
	repo.FileName = "UserRepositoryImpl.go"
	return &repo
}

func (repository UserRepositoryImpl) CreateUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel {
	funcName := "CreateUser"
	query := fmt.Sprintf(`INSERT INTO %s 
		(username, password, user_secret, first_name, last_name) 
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id
	`, repository.TableName)

	result := ctx.Tx.QueryRow(query, user.Username.String, user.Password.String,
		user.UserSecret.String, user.FirstName.String, user.LastName.String)
	err := result.Scan(&user.ID)

	helper.PanicIfErrorWithLocation(err, repository.FileName, funcName, ctx)
	return user
}

func (repository UserRepositoryImpl) UpdateUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) DeleteUser(ctx *applicationModel.ContextModel, user entity.UserModel) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) DetailUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) LoginUser(ctx *applicationModel.ContextModel, user entity.UserModel) (result entity.UserModel) {
	funcName := "LoginUser"
	query := fmt.Sprintf(
		`SELECT 
			id, username, password, user_secret
		FROM %s 
		WHERE
			username = $1`)

	rows, err := ctx.Tx.Query(query, user.Username.String)
	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Username, &result.Password, &result.UserSecret)
		if err != nil {
			helper.PanicIfError(err)
		}
	} else {
		helper.PanicIfError(errorModel.GenerateForbiddenAccessError(repository.FileName, funcName))
	}
	helper.PanicIfErrorWithLocation(err, repository.FileName, funcName, ctx)
	return
}

func (repository UserRepositoryImpl) GetListUser(ctx *applicationModel.ContextModel) []entity.UserModel {
	//TODO implement me
	panic("implement me")
}
