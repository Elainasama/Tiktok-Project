package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}

func Fail(c *gin.Context, msg string, v interface{}) {
	c.JSON(http.StatusOK)
}
