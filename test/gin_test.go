package test

import (
	
	"testing"
	"github.com/jyu/routers"
)

func TestGin(t *testing.T){
	r := routers.Router()
	r.Run()
}
