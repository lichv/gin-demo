package routers

import (
	"gin-demo/app/controllers/user"
	"gin-demo/app/middlewares"
	"gin-demo/utils/setting"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(favicon.New(path.Join(setting.AppSetting.RootPath, "favicon.ico")))

	r.Use(middlewares.Cors())

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"state":   2000,
			"message": "success",
		})
	})

	r.Any("/auth/login", user.Login)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.GET("/user/list", user.GetUsers)
	}

	return r
}
