package CategoryService

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/repository"
	"github.com/horcrux12/clean-rest-api-template/service"
	"net/http"
	"strconv"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
	service.AbstractService
}

func (service CategoryServiceImpl) ReadBodyAndParam(request *http.Request) (result in.CategoryUpdateRequest) {
	strBody := helper.ReadBody(request)
	if strBody != "" {
		errorS := json.Unmarshal([]byte(strBody), &result)
		helper.PanicIfError(errorS)
	}

	id, _ := strconv.Atoi(mux.Vars(request)["ID"])
	if result.ID == 0 {
		result.ID = int64(id)
	}

	return
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
		AbstractService: service.AbstractService{
			FileName:    "CategoryServiceImpl.go",
			ServiceName: "CATEGORY",
		},
	}
}
