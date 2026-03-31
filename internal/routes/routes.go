package routes

import (
	"net/http"

	"github.com/diaszakir/bookmark/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
