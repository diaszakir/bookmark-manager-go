package handlers

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/diaszakir/bookmark/internal/models"
	"github.com/gin-gonic/gin"
)

var bookmarks []models.Bookmark
var id int64

func GetBookmarks(c *gin.Context) {
	name := c.Query("name")
	tags := c.QueryArray("tag")

	if name == "" && len(tags) == 0 {
		c.JSON(http.StatusOK, bookmarks)
		return
	}

	var result []models.Bookmark

	for _, b := range bookmarks {
		if name != "" && b.Name == name {
			result = append(result, b)
			continue
		}

		if slices.Contains(tags, b.Tag) {
			result = append(result, b)
			break
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetBookmark(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse bookmark id"})
		return
	}

	for _, v := range bookmarks {
		if v.Id == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "bookmark not found"})
}

func CreateBookmark(c *gin.Context) {
	var bookmark models.Bookmark
	err := c.ShouldBindJSON(&bookmark)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	id++
	bookmark.Id = id

	bookmarks = append(bookmarks, bookmark)
	c.Status(http.StatusCreated)
}

func EditBookmark(c *gin.Context) {

}

func DeleteBookmark(c *gin.Context) {

}

func DeleteBookmarks(c *gin.Context) {

}
