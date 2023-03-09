package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	helpers "github.com/yogiis/golang-api-automation/app/crud-restful-api/helpers"
	model "github.com/yogiis/golang-api-automation/app/crud-restful-api/model"
	repository "github.com/yogiis/golang-api-automation/app/crud-restful-api/repository"
)

type UserService interface {
	Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse
	Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) model.UserResponse
	FindAll(ctx context.Context) []model.UserResponse
}

type NewUserService struct {
	UserRespository repository.UserRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func RouteUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &NewUserService{
		UserRespository: userRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (s *NewUserService) Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse {
	err := s.Validate.Struct(request)
	helpers.LogPanicError(err)

	tx, err := s.DB.Begin()
	helpers.LogPanicError(err)
	defer helpers.CommitOrRollback(tx)

	user := model.UserModel{
		Name: request.Name,
	}

	user = s.UserRespository.Save(ctx, tx, user)
	return ToUsersResponse(user)
}

func (s *NewUserService) Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse {
	err := s.Validate.Struct(request)
	helpers.LogPanicError(err)

	tx, err := s.DB.Begin()
	helpers.LogPanicError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRespository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(helpers.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name

	user = s.UserRespository.Update(ctx, tx, user)

	return ToUsersResponse(user)
}

func (s *NewUserService) Delete(ctx context.Context, userId int) {
	tx, err := s.DB.Begin()
	helpers.LogPanicError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRespository.FindById(ctx, tx, userId)
	if err != nil {
		panic(helpers.NewNotFoundError(err.Error()))
	}
	s.UserRespository.Delete(ctx, tx, user)
}

func (s *NewUserService) FindById(ctx context.Context, userId int) model.UserResponse {
	tx, err := s.DB.Begin()
	helpers.LogPanicError(err)
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRespository.FindById(ctx, tx, userId)
	if err != nil {
		panic(helpers.NewNotFoundError(err.Error()))
	}
	return ToUsersResponse(user)
}

func (s *NewUserService) FindAll(ctx context.Context) []model.UserResponse {
	tx, err := s.DB.Begin()
	helpers.LogPanicError(err)
	defer helpers.CommitOrRollback(tx)

	users := s.UserRespository.FindAll(ctx, tx)

	var usersResponses []model.UserResponse
	for _, user := range users {
		usersResponses = append(usersResponses, ToUsersResponse(user))
	}
	return usersResponses
}

func ToUsersResponse(user model.UserModel) model.UserResponse {
	return model.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}
}
