package main

import (
	"agenda-escolar/internal/http/handler"
	"agenda-escolar/internal/http/router"
	"os"
)

func main() {
	os.Setenv("production", "1")

	routes := router.NewRouter()
	handler.HandleRequests(routes)
	routes.Run(":8000")
}
