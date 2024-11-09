package controller

import (
	"agenda-escolar/internal/domain"
	"agenda-escolar/internal/services"
	"agenda-escolar/pkg"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vingarcia/ksql"
)

var taskService services.TaskService

type TaskController struct {
}

func (*TaskController) Create(c *gin.Context) {
	var (
		form  map[string]any
		input domain.CreateForm
		task  domain.Task
	)
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}
	if len(input.Date) == 10 {
		input.Date += " 00:00:00"
	}
	if !pkg.IsDateValid(input.Date) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": `invalid "date" field`})
		return
	}
	user_id := c.Keys["user_id"].(int)

	inputBytes, _ := json.Marshal(input)
	json.Unmarshal(inputBytes, &form)

	form["date"], _ = pkg.ParseDate(input.Date)
	form["userId"] = user_id

	formBytes, _ := json.Marshal(form)
	json.Unmarshal(formBytes, &task)

	taskCreated, err := taskService.Create(&task)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "cannot create task"})
		return
	}

	c.JSON(http.StatusCreated, taskCreated)
}

func (*TaskController) List(c *gin.Context) {
	dateFrom := c.Query("dateFrom")
	dateTo := c.Query("dateTo")
	if len(dateFrom) == 10 {
		dateFrom += " 00:00:00"
	}
	if len(dateTo) == 10 {
		dateTo += " 23:59:59"
	}
	if !pkg.IsDateValid(dateFrom) || !pkg.IsDateValid(dateTo) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "dateFrom or dateTo query param is invalid, requires format (yyyy-mm-dd)"})
		return
	}
	user_id := c.Keys["user_id"].(int)
	list, err := taskService.List(user_id, dateFrom, dateTo)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, list)
}

func (*TaskController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	task, err := taskService.GetByID(id)

	if err != nil {
		fmt.Println(err)
		statusCode := http.StatusInternalServerError
		if errors.Is(err, ksql.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
		}
		c.AbortWithStatusJSON(statusCode, gin.H{"message": "cannot get task"})
	}

	c.JSON(http.StatusOK, task)
}

func (*TaskController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid :id param"})
		return
	}
	form := make(map[string]any)
	err := c.BindJSON(&form)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	formBytes, _ := json.Marshal(form)
	var formData domain.Task

	json.Unmarshal(formBytes, &formData)
	formData.ID = &id

	if dt, ok := form["date"].(string); ok {
		formData.Date, _ = pkg.ParseDate(dt)
	}

	task, err := taskService.Update(&formData)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("cannot update: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, task)
}
