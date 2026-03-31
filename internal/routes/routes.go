package routes

import (
	"net/http"

	"github.com/diaszakir/bookmark/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/health", checkAPI)

	r.GET("/bookmark", handlers.GetBookmarks)
	r.GET("/bookmark/:id", handlers.GetBookmark)

	r.POST("/bookmark", handlers.CreateBookmark)

	r.PUT("/bookmark/:id", handlers.EditBookmark)

	r.DELETE("/bookmark/:id", handlers.DeleteBookmark)
	r.DELETE("/bookmark", handlers.DeleteBookmarks)
}

func checkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API Works"})
}
