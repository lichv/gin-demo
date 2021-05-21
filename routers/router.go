package routers

import (
	"gin-demo/app/controllers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r :=gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"state":2000,
			"message":"success",
		})
	})

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/user/list", user.GetUsers)
	}

	return r
}
