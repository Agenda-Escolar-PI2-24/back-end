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
	// user_id := c.Keys["user_id"].(int)

	// list, err := taskService.List(user_id)
	list, err := taskService.List()

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
	if err != nil || len(form) == 0 {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	var formData domain.Task

	formData.ID = &id
	if value, valid := form["title"].(string); valid {
		formData.Title = &value
	}
	if value, valid := form["obs"].(string); valid {
		formData.Obs = &value
	}
	if value, valid := form["content"].(string); valid {
		formData.Content = &value
	}
	if value, valid := form["class"].(string); valid {
		formData.Class = &value
	}
	if value, valid := form["userId"].(float64); valid {
		user_id := int(value)
		formData.UserID = &user_id
	}
	if value, valid := form["contempled"].(bool); valid {
		formData.Contempled = &value
	}
	if value, valid := form["satisfactory"].(bool); valid {
		formData.Satisfactory = &value
	}

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
