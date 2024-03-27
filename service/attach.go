package service

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jyu/utils"
)

func Upload(c *gin.Context){
	writer := c.Writer
	request := c.Request

	srcFile, head, err := request.FormFile("file")
	if err!= nil {
        utils.RespFailed(writer, err.Error())
    }

	suffix := ".png"
	ofileName := head.Filename
	tem := strings.Split(ofileName, ".")
	if len(tem) > 1{
		suffix = "." + tem[len(tem)-1]
	}

	filename := fmt.Sprintf("%d%04d%s", time.Now().Unix(),rand.Int31(),suffix)

	dstFile, err := os.Create("./Upload" + filename)
	if err!= nil {
        utils.RespFailed(writer, err.Error())
    }
	_, err = io.Copy(dstFile,srcFile)
	if err!= nil {
        utils.RespFailed(writer, err.Error())
    }

	url := "./Upload" + filename
	utils.RespOK(writer, url,"图片发送成功")
}

