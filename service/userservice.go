package service

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jyu/models"
	"github.com/jyu/utils"
)

// GetUserListHandle
// @Tags 用户模块
// @Success 200 {string} {"code", "masssge"}
// @Router /User/getUserList [GET]
func GetUserListHandler(c *gin.Context) {

	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"masssge": "查找成功",
		"data":    data,
	})
}

// RegisterHandler
// @Tags 用户模块
// @param account query string false "账户"
// @param password query string false "密码"
// @param repassword query string false "确认密码""
// @Success 200 {string} {"code", "masssge"}
// @Router /User/createUserHandler [POST]
func RegisterHandler(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": err.Error(),
			"data":    user,
		})
		return
	}
	if user.Account == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "账户和密码不能为空",
			"data":    user,
		})
		return
	}
	log.Println(user)
	data, _ := models.FindUserByAccount(user.Account)
	log.Println(data)
	if data.Account != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "账户已存在",
			"data":    nil,
		})
		return
	}

	user.Salt = fmt.Sprintf("%06", rand.Int31())
	user.Password = utils.MakePassword(user.Password, user.Salt)

	if user.LoginTime.IsZero() {
		user.LoginTime = time.Now()
	}
	if user.HeartbeatTime.IsZero() {
		user.HeartbeatTime = time.Now()
	}
	if user.LoginOutTime.IsZero() {
		user.LoginOutTime = time.Now()
	}
	db := models.CreateUser(user)
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "新增用户失败",
			"data":    nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"masssge": "新增用户成功",
			"data":    nil,
		})
	}
}

// DeleteUserHandler
// @Tags 删除模块
// @param account query string false "账户"
// @param password query string false "密码"
// @Success 200 {string} {"code", "masssge"}
// @Router /User/deleteUserHandler [GET]
func LogoutHandler(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": err.Error(),
			"data":    user,
		})
		return
	}
	data, _ := models.FindUserByAccount(user.Account)
	log.Println(data)
	if data.Account != "" {
		db := models.DeleteUser(user)
		if db.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"masssge": "删除用户失败",
				"data":    nil,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"masssge": "删除用户成功",
			"data":    nil,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "账户不存在",
			"data":    data,
		})
	}

}

func UpdateUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": err.Error(),
			"data":    user,
		})
		return
	}
	log.Println(user)
	db := models.UpdateUser(user)
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "更新用户失败",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"masssge": "更新成功",
		"data":    nil,
	})
}

func LoginHandler(c *gin.Context) {
	//接收请求
	user := models.UserBasic{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": err.Error(),
			"data":    user,
		})
		return
	}
	//找用户
	data, _ := models.FindUserByAccount(user.Account)
	if data.Account == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "账户不存在",
			"data":    data,
		})
		return
	}

	// 密码判断
	flag := utils.ValidPassword(user.Password, data.Salt, data.Password)

	// 更新token
	db := models.UpdateIdentity(data)
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "token更新失败",
			"data":    data,
		})
		return
	}

	if flag {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"masssge": "登录成功",
			"data":    data,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    0,
			"masssge": "密码错误",
			"data":    data,
		})
	}
}

var upGrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			log.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}

// 死循环 噢
func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			log.Println(err)
			return
		}

		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:[%s]", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	log.Println(c.Query("account"))
	data, err := models.SearchFriend(c.Query("account"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"masssge": "查询失败",
			"data":    nil,
		})
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"code":    0,
	// 	"masssge": "查询成功",
	// 	"data":    data,
	// })
	utils.RespOKList(c.Writer,data,len(data))
}
