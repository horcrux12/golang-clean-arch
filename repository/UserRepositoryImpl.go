package repository

import (
	"database/sql"
	"fmt"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
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

func (repository UserRepositoryImpl) CreateUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel {
	funcName := "CreateUser"
	query := fmt.Sprintf(`INSERT INTO %s 
		(username, password, user_secret, first_name, last_name) 
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id
	`, repository.TableName)

	result := tx.QueryRow(query, user.Username.String, user.Password.String,
		user.UserSecret.String, user.FirstName.String, user.LastName.String)
	err := result.Scan(&user.ID)

	helper.PanicIfErrorWithLocation(err, repository.FileName, funcName, ctx)
	return user
}

func (repository UserRepositoryImpl) UpdateUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) DeleteUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) DetailUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) LoginUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel {
	//TODO implement me
	panic("implement me")
}

func (repository UserRepositoryImpl) GetListUser(ctx *applicationModel.ContextModel, tx *sql.Tx) []entity.UserModel {
	//TODO implement me
	panic("implement me")
}
