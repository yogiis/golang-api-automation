package model

type UserCreateRequest struct {
	Name string `validate:"required,max=200,min=1" json:"name" form:"name"`
}

type UserUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}
