package repository

import (
	"context"
	"database/sql"
	"errors"

	helpers "github.com/yogiis/golang-api-automation/app/crud-restful-api/helpers"

	model "github.com/yogiis/golang-api-automation/app/crud-restful-api/model"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user model.UserModel) model.UserModel
	Update(ctx context.Context, tx *sql.Tx, user model.UserModel) model.UserModel
	Delete(ctx context.Context, tx *sql.Tx, user model.UserModel)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (model.UserModel, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.UserModel
}

type NewUserRepository struct {
}

func RouteUserRepository() UserRepository {
	return &NewUserRepository{}
}

func (repository *NewUserRepository) Save(ctx context.Context, tx *sql.Tx, user model.UserModel) model.UserModel {
	SQL := "insert into users(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name)
	helpers.LogPanicError(err)

	id, err := result.LastInsertId()
	helpers.LogPanicError(err)

	user.Id = int(id)

	return user
}

func (repository *NewUserRepository) Update(ctx context.Context, tx *sql.Tx, user model.UserModel) model.UserModel {
	SQL := "update users set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Id)
	helpers.LogPanicError(err)

	return user
}

func (repository *NewUserRepository) Delete(ctx context.Context, tx *sql.Tx, user model.UserModel) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helpers.LogPanicError(err)
}

func (repository *NewUserRepository) FindById(ctx context.Context, tx *sql.Tx, userId int) (model.UserModel, error) {
	SQL := "select id, name from users where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helpers.LogPanicError(err)
	defer rows.Close()

	user := model.UserModel{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		helpers.LogPanicError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *NewUserRepository) FindAll(ctx context.Context, tx *sql.Tx) []model.UserModel {
	SQL := "select id, name from users"
	rows, err := tx.QueryContext(ctx, SQL)
	helpers.LogPanicError(err)
	defer rows.Close()

	var users []model.UserModel

	for rows.Next() {
		user := model.UserModel{}
		err := rows.Scan(&user.Id, &user.Name)
		helpers.LogPanicError(err)
		users = append(users, user)
	}
	return users
}
