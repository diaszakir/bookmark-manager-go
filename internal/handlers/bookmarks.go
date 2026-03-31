package handlers

import (
	"cmp"
	"net/http"
	"slices"
	"strconv"

	"github.com/diaszakir/bookmark/internal/models"
	"github.com/gin-gonic/gin"
)

var bookmarks []models.Bookmark
var id int64 = 0

// GetBookmarks godoc
// @Summary     List bookmarks
// @Tags        bookmarks
// @Produce     json
// @Param       name  query     string  false  "Filter by name"
// @Param       tag   query     []string false  "Filter by tag"
// @Success     200   {array}   models.Bookmark
// @Router      /bookmark [get]
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

// GetBookmark godoc
// @Summary     Get a bookmark
// @Tags        bookmarks
// @Produce     json
// @Param       id   path      int  true  "Bookmark ID"
// @Success     200  {object}  models.Bookmark
// @Failure     400  {object}  map[string]string
// @Failure     404  {object}  map[string]string
// @Router      /bookmark/{id} [get]
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

// CreateBookmark godoc
// @Summary     Create a bookmark
// @Tags        bookmarks
// @Accept      json
// @Param       body  body      models.Bookmark  true  "Bookmark data"
// @Success     201
// @Failure     400  {object}  map[string]string
// @Router      /bookmark [post]
func CreateBookmark(c *gin.Context) {
	var bookmark models.Bookmark
	err := c.ShouldBindJSON(&bookmark)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id++
	bookmark.Id = id

	bookmarks = append(bookmarks, bookmark)
	c.Status(http.StatusCreated)
}

// EditBookmark godoc
// @Summary     Update a bookmark
// @Tags        bookmarks
// @Accept      json
// @Param       id    path      int              true  "Bookmark ID"
// @Param       body  body      models.Bookmark  true  "Bookmark data"
// @Success     204
// @Failure     400  {object}  map[string]string
// @Failure     404  {object}  map[string]string
// @Router      /bookmark/{id} [put]
func EditBookmark(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse bookmark id"})
		return
	}

	var editBookmark models.Bookmark
	err = c.ShouldBindJSON(&editBookmark)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	editBookmark.Id = id

	for i, v := range bookmarks {
		if v.Id == id {
			bookmarks[i].Name = cmp.Or(editBookmark.Name, bookmarks[i].Name)
			bookmarks[i].URL = cmp.Or(editBookmark.URL, bookmarks[i].URL)
			bookmarks[i].Tag = cmp.Or(editBookmark.Tag, bookmarks[i].Tag)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "bookmark not found"})
}

// DeleteBookmark godoc
// @Summary     Delete a bookmark
// @Tags        bookmarks
// @Param       id   path      int  true  "Bookmark ID"
// @Success     204
// @Failure     400  {object}  map[string]string
// @Router      /bookmark/{id} [delete]
func DeleteBookmark(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse bookmark id"})
		return
	}

	bookmarks = slices.Delete(bookmarks, int(id-1), int(id))
	c.Status(http.StatusNoContent)
}

// DeleteBookmarks godoc
// @Summary     Delete all bookmarks
// @Tags        bookmarks
// @Success     204
// @Router      /bookmark [delete]
func DeleteBookmarks(c *gin.Context) {
	bookmarks = nil
	c.Status(http.StatusNoContent)
}
