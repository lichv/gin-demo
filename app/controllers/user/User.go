package user

import (
	"fmt"
	"gin-demo/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users,errs:=services.GetWhitelistUsers(map[string]interface{}{},"code desc",-1)
	fmt.Println(users)
	fmt.Println(errs)
	if len(errs) > 0 {
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
	})
}

