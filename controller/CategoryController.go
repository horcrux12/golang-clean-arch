package controller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/repository/postgresql"
	"github.com/horcrux12/clean-rest-api-template/service/CategoryService"
	"net/http"
	"strconv"
)

type CategoryController struct {
	CategoryService CategoryService.CategoryService
	AbstractController
}

//var CategoryController = CategoryController{}.New()

func NewCategoryController(validate *validator.Validate) CategoryController {
	categoryRepository := postgresql.NewCategoryRepository()
	return CategoryController{
		CategoryService: CategoryService.NewCategoryService(categoryRepository, validate),
		AbstractController: AbstractController{
			FileName: "CategoryController.go",
		},
	}
}

func (controller CategoryController) CategoryControllerWithoutPathParam(writer http.ResponseWriter, request *http.Request) {
	funcName := "CategoryControllerWithoutPathParam"
	var contextModel *applicationModel.ContextModel
	var payload out.WebResponse

	switch request.Method {
	case "POST":
		var inputStruct in.CategoryRequest
		helper.ReadFromRequestBody(request, &inputStruct)

		contextModel = controller.WhiteListServe(funcName, writer, request)
		payload = controller.CategoryService.Create(contextModel, inputStruct)
		break
	case "GET":
		contextModel = controller.WhiteListServe(funcName, writer, request)
		payload = controller.CategoryService.FindAll(contextModel)
		break
	}
	defer controller.LogResponse(contextModel, payload, funcName)
	helper.WriteToResponseBody(writer, payload)
}

func (controller CategoryController) CategoryControllerWithPathParam(writer http.ResponseWriter, request *http.Request) {
	funcName := "CategoryControllerWithPathParam"
	var contextModel *applicationModel.ContextModel
	var payload out.WebResponse
	var inputRequest in.CategoryRequest

	inputRequest = controller.readBodyAndParam(request)

	switch request.Method {
	case http.MethodGet:
		contextModel = controller.WhiteListServe(funcName, writer, request)
		payload = controller.CategoryService.FindByID(contextModel, inputRequest)
		break
	case http.MethodPut:
		inputRequest.IsUpdate = true
		contextModel = controller.WhiteListServe(funcName, writer, request)
		payload = controller.CategoryService.Update(contextModel, inputRequest)
		break
	case http.MethodDelete:
		inputRequest.IsDelete = true
		contextModel = controller.WhiteListServe(funcName, writer, request)
		payload = controller.CategoryService.Delete(contextModel, inputRequest)
	}
	defer controller.LogResponse(contextModel, payload, funcName)
	helper.WriteToResponseBody(writer, payload)
}

func (controller CategoryController) readBodyAndParam(request *http.Request) (result in.CategoryRequest) {
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
