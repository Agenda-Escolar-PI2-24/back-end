package handler

import (
	"agenda-escolar/internal/http/controller"
	"agenda-escolar/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var agendaController controller.TaskController

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func HandleRequests(router *gin.Engine) {
	router.GET("/", healthCheck)
	v1 := router.Group("/v1")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/login", controller.Login)
		authRoutes.POST("/register", controller.Register)
	}

	//requires authentication
	agendaRoutes := v1.Group("/agenda")
	agendaRoutes.Use(pkg.AuthenticationMiddleware())
	{
		agendaRoutes.POST("", agendaController.Create)
		agendaRoutes.GET("", agendaController.List)
		agendaRoutes.GET("/:id", agendaController.GetByID)
		agendaRoutes.PUT("/:id", agendaController.Update)
	}

}
