package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jyu/models"
)

// GetIndexHandle
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndexHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"masssge": "welcome",
	})
}

func Chat(c*gin.Context){
	models.Chat(c.Writer, c.Request)
}