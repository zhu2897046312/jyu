package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strings"
)

//转小写
func Md5EnCode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}
//转大写
func MD5EnCode(s string) string {
	return strings.ToUpper(Md5EnCode(s))
}

//随机数加密
func MakePassword(plainpwd string, salt string) string {
	return Md5EnCode(plainpwd + salt)
}

//随机数解密
func ValidPassword(plainpwd string, salt string, password string) (bool){
	md := Md5EnCode(plainpwd + salt)
	log.Println(md + "-----------" + password)
	return  md  == password
}