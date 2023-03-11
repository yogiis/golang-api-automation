package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/controller"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/helpers"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/middleware"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/repository"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/service"
)

func main() {

	db := helpers.NewDB()
	validate := validator.New()
	userRespository := repository.RouteUserRepository()
	userService := service.RouteUserService(userRespository, db, validate)
	userController := controller.RouteUserController(userService)

	router := httprouter.New()

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId", userController.FindById)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	router.PanicHandler = helpers.ErrorHandler

	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helpers.LogPanicError(err)
}
