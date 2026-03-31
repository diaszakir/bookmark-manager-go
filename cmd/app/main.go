package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/diaszakir/bookmark/docs"
	"github.com/diaszakir/bookmark/internal/routes"
)

// @title           Bookmark API
// @version         1.0
// @host            localhost:8080
// @BasePath        /
func main() {
	app := gin.Default()

	routes.Routes(app)

	app.Run(":8080")
}
