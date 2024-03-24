package main

import (
	"github.com/jyu/routers"
	"github.com/jyu/utils"
)

func main(){
	utils.InitConfig()
	utils.InitMySQL()
	r := routers.Router()
	r.Run()
}