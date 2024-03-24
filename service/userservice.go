package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jyu/models"
)

// GetUserListHandle
// @Tags 用户模块
// @Success 200 {string} {"code", "masssge"}
// @Router /User/getUserList [get]
func GetUserListHandler(c *gin.Context) {

	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(http.StatusOK, gin.H{
		"masssge": data,
	})
}

// CreateUserHandler
// @Tags 用户模块
// @param account query string false "账户"
// @param password query string false "密码"
// @param repassword query string false "确认密码""
// @Success 200 {string} {"code", "masssge"}
// @Router /User/createUserHandler [POST]
func CreateUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	// user.Account = c.Query("account")
	// user.Password = c.Query("password")
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"masssge": err.Error(),
		})
		return
	}
	log.Println(user)

	data, _ := models.FindUserByAccount(user.Account)
	log.Println(data)
	if  data.Account != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"masssge": "账户已存在",
		})
	    return
	}
	if user.LoginTime.IsZero() {
		user.LoginTime = time.Now()
	}
	if user.HeartbeatTime.IsZero() {
		user.HeartbeatTime = time.Now()
	}
	if user.LoginOutTime.IsZero() {
		user.LoginOutTime = time.Now()
	}
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"masssge": "新增用户成功",
	})

}

// DeleteUserHandler
// @Tags 删除模块
// @param account query string false "账户"
// @param password query string false "密码"
// @Success 200 {string} {"code", "masssge"}
// @Router /User/deleteUserHandler [get]
func DeleteUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"masssge": err.Error(),
		})
		return
	}
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"masssge": "删除用户成功",
	})
}

func UpdateUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"masssge": err.Error(),
		})
		return
	}
	log.Println(user)
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"masssge": "更新成功",
	})
}
