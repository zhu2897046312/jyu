package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FormID      string //发送者
	TargetID    string //接收者
	Type        int    //消息类型 群聊 、 私聊 、 广播
	Media       int    //消息类型 文字、图片、音频
	Context     string //消息内容
	Pic         string
	Url         string
	Description string
	Amount      int
}

func (table *Message) TableNanme() string {
	return "message"
}
