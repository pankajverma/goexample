package routergrouping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"stash.edgewaternetworks.com/scc/gin_prj/modal"
)

/*
CreateRouteGroup ...
*/
func CreateRouteGroup(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/group1", group1Route)
		v1.GET("/hello", group1RouteHello)
	}

	v2 := router.Group("v2")
	{
		v2.GET("/group2", group1Route)
		v2.GET("/hello", group1RouteHello)
	}
}

func group1Route(c *gin.Context) {
	c.JSON(http.StatusOK, modal.Success{Status: 200, Message: "this router v1/group1 works."})
}

func group1RouteHello(c *gin.Context) {
	c.JSON(http.StatusOK, modal.Success{Status: 200, Message: "Say hello to route v1: group1RouteHello"})
}

func group2Route(c *gin.Context) {
	c.JSON(http.StatusOK, modal.Success{Status: 200, Message: "this router v2/group1 works."})
}

func group2RouteHello(c *gin.Context) {
	c.JSON(http.StatusOK, modal.Success{Status: 200, Message: "Say hello to route v2: group2RouteHello"})
}
