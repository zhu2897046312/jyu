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
	r.GET("/User/LogoutHandler",service.LogoutHandler)
	r.GET("/User/Chat",service.Chat)
	r.GET("/User/SendMsg",service.SendMsg)
	r.GET("/User/SendUserMsg",service.SendUserMsg)
	r.GET("/User/SearchFriend",service.SearchFriends)
	r.GET("/User/LoadCommunityList",service.LoadCommunityList)


	r.POST("/User/UpdateUserHandler",service.UpdateUserHandler)
	r.POST("/User/RegisterHandler",service.RegisterHandler)
	r.POST("/User/LoginHandler",service.LoginHandler)
	r.POST("/attach/Upload",service.Upload)
	r.POST("/User/AddFriend",service.AddFriend)
	r.POST("/User/CreateCommunity",service.CreateCommunity)
	
	return r
}