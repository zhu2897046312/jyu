package test

import (
	"testing"
	"github.com/jyu/utils"
)
func Test_config(t *testing.T){
	utils.InitConfig()
	utils.InitMySQL()
}