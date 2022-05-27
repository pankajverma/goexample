package modalbinding

import (
	"fmt"
	"log"
	"net/http"

	"modal"

	"github.com/gin-gonic/gin"
)

/*
GenericBind ...
*/
func GenericBind(router *gin.Engine) {
	router.POST("/jsonBind", func(c *gin.Context) {
		var loginJSON modal.Login
		// ShouldBind is generic method which detect the content
		// based on header and use the appropoate bining engine
		if err := c.ShouldBind(&loginJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			msg := fmt.Sprintf("you tried with user:%s password:%s", loginJSON.User, loginJSON.Password)
			log.Print(msg)
			c.JSON(http.StatusOK, gin.H{"Status": msg})
		}
	})
}

/*
UserBind ...
*/
func UserBind(router *gin.Engine) {
	router.POST("/userBind", func(c *gin.Context) {
		var userType modal.User
		if err := c.ShouldBind(&userType); err != nil {
			errorMessage := modal.ErrorMessage{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			}
			c.JSON(http.StatusBadRequest, errorMessage)
		} else {
			success := modal.Success{
				Message: "this was good request" + userType.Name,
				Status:  http.StatusOK,
			}
			c.JSON(http.StatusOK, success)
		}
	})
}
