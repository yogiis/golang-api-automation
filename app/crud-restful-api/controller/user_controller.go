package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	helpers "github.com/yogiis/golang-api-automation/app/crud-restful-api/helpers"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/model"
	"github.com/yogiis/golang-api-automation/app/crud-restful-api/service"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type NewUserController struct {
	UserService service.UserService
}

func RouteUserController(userService service.UserService) UserController {
	return &NewUserController{
		UserService: userService,
	}
}

func (c *NewUserController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	request.ParseForm()
	userCreateRequest := model.UserCreateRequest{
		Name: request.FormValue("name"),
	}

	userResponse := c.UserService.Create(request.Context(), userCreateRequest)
	newUserResponse := model.FormatterResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helpers.WriteToResponseBody(writer, newUserResponse)
}

func (c *NewUserController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	userUpdateRequest := model.UserUpdateRequest{}
	helpers.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helpers.LogPanicError(err)

	userUpdateRequest.Id = id

	userResponse := c.UserService.Update(request.Context(), userUpdateRequest)
	newUserResponse := model.FormatterResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helpers.WriteToResponseBody(writer, newUserResponse)
}

func (c *NewUserController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helpers.LogPanicError(err)

	c.UserService.Delete(request.Context(), id)
	newUserResponse := model.FormatterResponse{
		Code:   200,
		Status: "OK",
	}

	helpers.WriteToResponseBody(writer, newUserResponse)
}

func (c *NewUserController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helpers.LogPanicError(err)

	userResponse := c.UserService.FindById(request.Context(), id)
	newUserResponse := model.FormatterResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helpers.WriteToResponseBody(writer, newUserResponse)
}

func (c *NewUserController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersResponses := c.UserService.FindAll(request.Context())
	newUserResponse := model.FormatterResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResponses,
	}

	helpers.WriteToResponseBody(writer, newUserResponse)
}
