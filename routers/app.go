package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jyu/service"

	docs "github.com/jyu/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine{
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index",service.GetIndexHandler)
	r.GET("/User/getUserList",service.GetUserListHandler)
	r.POST("/User/RegisterHandler",service.RegisterHandler)
	r.GET("/User/LogoutHandler",service.LogoutHandler)
	r.POST("/User/UpdateUserHandler",service.UpdateUserHandler)
	r.POST("/User/LoginHandler",service.LoginHandler)
	return r
}