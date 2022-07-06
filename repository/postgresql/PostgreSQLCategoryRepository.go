package postgresql

import (
	"errors"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/repository"
)

type PostgreSQLCategoryRepository struct {
	repository.AbstractRepository
}

func NewCategoryRepository() repository.CategoryRepository {
	repo := PostgreSQLCategoryRepository{}
	repo.TableName = "category"
	return &repo
}

func (repository PostgreSQLCategoryRepository) Save(ctx *applicationModel.ContextModel, category entity.CategoryModel) entity.CategoryModel {
	query := "INSERT into " + repository.TableName + " (name) VALUES ($1) RETURNING id"

	result := ctx.ConnectionModel.Tx.QueryRow(query, category.Name.String)
	err := result.Scan(&category.ID)

	helper.PanicIfError(err)
	return category
}

func (repository PostgreSQLCategoryRepository) Update(ctx *applicationModel.ContextModel, category entity.CategoryModel) entity.CategoryModel {
	funcName := "Update"
	query := "UPDATE " + repository.TableName + " SET name = $1 WHERE id = $2"

	_, err := ctx.ConnectionModel.Tx.Exec(query, category.Name.String, category.ID.Int64)
	helper.PanicIfErrorWithLocation(err, repository.FileName, funcName, ctx)

	return category
}

func (repository PostgreSQLCategoryRepository) Delete(ctx *applicationModel.ContextModel, category entity.CategoryModel) {
	query := "DELETE FROM " + repository.TableName + " WHERE id = $1"

	_, err := ctx.ConnectionModel.Tx.Exec(query, category.ID.Int64)
	helper.PanicIfError(err)
}

func (repository PostgreSQLCategoryRepository) FindByID(ctx *applicationModel.ContextModel, category entity.CategoryModel) (result entity.CategoryModel, err error) {
	query := "SELECT id, name FROM " + repository.TableName + " WHERE id = $1"

	rows, err := ctx.ConnectionModel.Tx.Query(query, category.ID.Int64)
	defer rows.Close()
	helper.PanicIfError(err)

	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		helper.PanicIfError(err)
		return
	} else {
		err = errors.New("category is not found")
		return
	}
	return
}

func (repository PostgreSQLCategoryRepository) FindAll(ctx *applicationModel.ContextModel) (result []entity.CategoryModel) {
	query := "SELECT id, name FROM " + repository.TableName
	rows, err := ctx.ConnectionModel.Tx.Query(query)
	defer rows.Close()
	helper.PanicIfError(err)

	for rows.Next() {
		var tempCategory entity.CategoryModel
		err = rows.Scan(&tempCategory.ID, &tempCategory.Name)
		helper.PanicIfError(err)
		result = append(result, tempCategory)
	}

	return
}
