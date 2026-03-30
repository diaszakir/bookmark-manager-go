package main

import (
	"github.com/gin-gonic/gin"

	"github.com/diaszakir/bookmark/internal/routes"
)

func main() {
	app := gin.Default()

	routes.Routes(app)

	app.Run(":8080")
}
