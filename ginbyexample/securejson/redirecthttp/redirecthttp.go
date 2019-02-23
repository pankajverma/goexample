package redirecthttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
RedirectTest ...
*/
func RedirectTest(router *gin.Engine) {
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "http://google.com")
	})
}

/*
RouterRedirectDemo ...
*/
func RouterRedirectDemo(router *gin.Engine) {
	router.GET("/demo", func(c *gin.Context) {
		c.Request.URL.Path = "/ping"
		router.HandleContext(c)
	})
}
