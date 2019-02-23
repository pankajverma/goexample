package securejson

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
GetSecureJSON ...
*/
func GetSecureJSON(router *gin.Engine) {
	router.GET("/getSecureJSON", func(c *gin.Context) {
		names := []string{"hello", "this", "is", "it"}
		c.SecureJSON(http.StatusOK, names)
	})
}
