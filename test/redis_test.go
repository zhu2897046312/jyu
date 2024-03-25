package test

import (
	"testing"
	"github.com/jyu/utils"
)
func Test_Redis(t *testing.T){
	utils.InitConfig("../config")
	utils.InitRedis()
}