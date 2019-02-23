package main

import (
	"net/http"

	"stash.edgewaternetworks.com/scc/gin_prj/redirecthttp"
	"stash.edgewaternetworks.com/scc/gin_prj/routergrouping"
	"stash.edgewaternetworks.com/scc/gin_prj/securejson"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"stash.edgewaternetworks.com/scc/gin_prj/customevalidator"
	"stash.edgewaternetworks.com/scc/gin_prj/modalbinding"

	"github.com/gin-gonic/gin"
)

func main() {

	v := binding.Validator.Engine().(*validator.Validate)
	customevalidator.RegisterCustomeValidators(v)

	router := gin.Default()
	baseGET(router)
	paramGET(router)
	queryParamGET(router)
	formPOST(router)

	modalbinding.GenericBind(router)
	modalbinding.UserBind(router)

	securejson.GetSecureJSON(router)

	redirecthttp.RouterRedirectDemo(router)
	redirecthttp.RedirectTest(router)

	routergrouping.CreateRouteGroup(router)
	// router.Run()
	router.Run(":3030") // to run on different port 3030
}

func baseGET(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func paramGET(router *gin.Engine) {
	router.GET("/ping/:name", func(c *gin.Context) {
		nameParam := c.Param("name")

		c.String(http.StatusOK, "Hello, name is: %s", nameParam)
	})

	router.GET("/ping/:name/*action", func(c *gin.Context) {
		nameParam := c.Param("name")
		action := c.Param("action")

		c.String(http.StatusOK, "Hello, name is: %s, action is: %s", nameParam, action)
	})
}

func queryParamGET(router *gin.Engine) {
	// "/ping?name=abc&age=18"
	router.GET("/details", func(c *gin.Context) {
		name := c.DefaultQuery("name", "guest")
		age := c.Query("age")

		c.String(http.StatusOK, "hello ur name is: %s and age is %s", name, age)

	})
}

func formPOST(router *gin.Engine) {
	router.POST("/post", func(c *gin.Context) {
		msg := c.PostForm("msg")
		name := c.PostForm("name")
		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"msg":    msg,
			"name":   name,
		})
	})
}
