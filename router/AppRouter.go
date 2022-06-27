package router

import (
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/controller"
)

var AppRouter appRouters

type appRouters struct {
	CategoryController controller.CategoryController
	UserController     controller.UserController
}

func InitiateRouter() {
	AppRouter.CategoryController = controller.NewCategoryController(app.ApplicationAttribute.DBConnection, app.ApplicationAttribute.Validate)
	AppRouter.UserController = controller.NewUserController(app.ApplicationAttribute.DBConnection, app.ApplicationAttribute.Validate)
}
