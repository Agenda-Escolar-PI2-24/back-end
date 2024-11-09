package controller

import (
	"agenda-escolar/internal/domain"
	"agenda-escolar/internal/services"
	"agenda-escolar/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService services.UserService

// Function for logging in
func Login(c *gin.Context) {
	var user domain.User

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Check if credentials are valid (replace this logic with real authentication)
	if user, err := userService.Login(&user); err == nil {

		// Generate a JWT token
		token, err := pkg.GenerateToken(int(user.ID), user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// Function for registering a new user (for demonstration purposes)
func Register(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Remember to securely hash passwords before storing them
	err := userService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("cannot register user: %s", err.Error())})
		return
	}
	//user.ID = 1 // Just for demonstration purposes
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
