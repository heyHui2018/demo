package routers

import (
	"github.com/heyHui2018/demo/gin/reserveProgram/middleWare"
	"github.com/heyHui2018/demo/gin/reserveProgram/controller"
	"github.com/heyHui2018/demo/gin/reserveProgram/base"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(base.GetConfig().Server.RunMode)

	pre := r.Group("/reserveProgram")
	pre.Use(middleWare.GenerateProcessId())
	{
		bind := pre.Group("/bind")
		{
			bind.GET("/query", controller.Query)
		}
	}
	return r
}
