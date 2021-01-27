package handler

import (
	"go-restapi/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewsfeedGet endpoint
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
