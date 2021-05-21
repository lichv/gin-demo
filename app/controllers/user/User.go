package user

import (
	"gin-demo/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users,total,errs := services.GetUserPages(map[string]interface{}{},"code desc",0,20)
	if errs != nil {
		c.JSON(http.StatusOK,gin.H{
			"state":3000,
			"message":"error",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"state":2000,
		"message":"success",
		"data":users,
		"total":total,
	})
}


