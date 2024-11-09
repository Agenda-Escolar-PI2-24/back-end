package repository

import (
	"agenda-escolar/internal/domain"
	"agenda-escolar/internal/storage/database"
	"context"
	"fmt"
)

type UserRepository struct {
}

func (*UserRepository) Register(user domain.User) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()

	qry := `insert into "user" (username, password) values ($1,$2)`
	_, err = db.Exec(ctx, qry, user.Username, user.Password)

	return
}

func (*UserRepository) Auth(user *domain.User) (domain.User, error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()

	var data []domain.User
	qry := `select _id, username from "user" where username = $1 and password = $2`
	err := db.Query(ctx, &data, qry, user.Username, user.Password)
	if err != nil {
		return domain.User{}, err
	}
	if len(data) == 0 {
		return domain.User{}, fmt.Errorf("user not found")
	}

	return data[0], nil
}

func (*UserRepository) Exists(username string) (bool, error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()

	var data []domain.User
	qry := `select username from "user" where username = $1`
	err := db.Query(ctx, &data, qry, username)
	if err != nil {
		return false, err
	}

	return len(data) > 0, nil
}
