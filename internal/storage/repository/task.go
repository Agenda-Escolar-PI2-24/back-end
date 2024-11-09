package repository

import (
	"agenda-escolar/internal/domain"
	"agenda-escolar/internal/storage/database"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type TaskRepository struct {
}

func (*TaskRepository) Create(task *domain.Task) (*domain.Task, error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()

	qry := "insert into task (title, class, content, date, user_id) values ($1, $2, $3, $4, $5) returning *"
	err := db.QueryOne(ctx, task, qry, task.Title, task.Class, task.Content, task.Date, task.UserID)

	return task, err
}

func (*TaskRepository) List(userID int, dateFrom, dateTo string) (list []domain.Task, err error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()
	list = make([]domain.Task, 0)

	err = db.Query(ctx, &list, "select * from task where user_id = $1 and date between $2 and $3", userID, dateFrom, dateTo)

	return
}

func (tr *TaskRepository) Update(task *domain.Task) (*domain.Task, error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()

	var updateParams []any
	params, fields := getParamsAndFields(*task, false)
	updateParams = append(updateParams, task.ID)
	updateParams = append(updateParams, params...)

	qry := tr.getUpdateQuery(fields)

	err := db.QueryOne(ctx, task, qry, updateParams...)

	return task, err
}

func (*TaskRepository) GetByID(id int) (task *domain.Task, err error) {
	db := database.GetDB()
	ctx := context.Background()
	defer db.Close()
	var t domain.Task

	err = db.QueryOne(ctx, &t, "select * from task where _id = $1", id)
	task = &t

	return
}

func getParamsAndFields(data domain.Task, isCreation bool) (params []any, fields []string) {
	v := reflect.ValueOf(data)

	//v = v.Elem()
	tipo := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := strings.ToLower(tipo.Field(i).Name)
		if !isCreation {
			if tipo.Field(i).Tag.Get("ksql") == "" || fieldName == "id" {
				continue
			}
		}
		if tipo.Field(i).Tag.Get("ksql") == "" || fieldName == "id" {
			continue
		}

		if field.Kind() == reflect.Ptr {
			if !field.IsNil() {
				if field.Elem().Kind() == reflect.Map || field.Elem().Kind() == reflect.Array {
					jsonBytes, err := json.Marshal(field.Elem().Interface())
					if err != nil {
						fmt.Println("cannot serialize", fieldName)
						continue
					}
					params = append(params, jsonBytes)
				} else {
					params = append(params, field.Elem().Interface())
				}
				fields = append(fields, tipo.Field(i).Tag.Get("ksql"))
			}
			continue
		}
		params = append(params, field.Interface())
		fields = append(fields, tipo.Field(i).Tag.Get("ksql"))
	}

	return
}

func (*TaskRepository) getUpdateQuery(fields []string) string {
	var params []string
	for i, field := range fields {
		params = append(params, fmt.Sprintf("%s = $%d\n", field, i+2))
	}

	return fmt.Sprintf(`
		UPDATE task SET 
			%s
		WHERE _id = $1
		RETURNING *
	`, strings.Join(params, ","))
}
