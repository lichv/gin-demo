package routers

import (
	Home "gin-demo/app/controllers/home"
	"gin-demo/app/controllers/user"
	"gin-demo/app/middlewares"
	"gin-demo/utils/setting"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"path"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(favicon.New(path.Join(setting.AppSetting.RootPath, "favicon.ico")))
	r.LoadHTMLGlob("./public/*.html")
	r.Static("/static", "./public/static/")
	r.Use(middlewares.Cors())

	r.GET("/", Home.Index)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewares.JWT())
	{
		apiv1.GET("/user/list", user.GetUsers)
	}

	return r
}
