package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/health", checkAPI)

	r.GET("/bookmark")
	r.GET("/bookmark/:id")

	r.POST("/bookmark")

	r.PUT("/bookmark/:id")

	r.DELETE("/bookmark/:id")
}

func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API Works"})
}
