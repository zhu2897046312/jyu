package utils

import (
	"log"
	"net/http"
	"encoding/json"
)

type H struct{
	Code int 
	Msg string
	Data interface{}
	Rows interface{}
	Total interface{}
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Data: data,
		Msg: msg,
	}
	ret, err := json.Marshal(h)
	if err!= nil {
       log.Println(err)
        return
    }

	w.Write(ret)
}

func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Rows: data,
		Total: total,
	}
	ret, err := json.Marshal(h)
	if err!= nil {
       log.Println(err)
        return
    }

	w.Write(ret)
}

func RespFailed(w http.ResponseWriter, msg string){
	Resp(w, 0, nil, msg)
}

func RespOK(w http.ResponseWriter, data interface{}, msg string){
    Resp(w, 0, data, "")
}

func RespOKList(w http.ResponseWriter, data interface{}, total interface{}){
	RespList(w, 0, data, total)
}

