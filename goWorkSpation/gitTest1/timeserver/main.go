package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type currentDate struct {
	Current string `json:"current"`
}

func main() {

	//处理请求
	http.HandleFunc("/go", func(rw http.ResponseWriter, request *http.Request) {
		rs := &currentDate{}
		defer func() { OutputJson(rw, rs) }()

		cur := time.Now().Format("2006-01-02 15:04:05")
		rs.Current = cur

	})
	http.ListenAndServe("127.0.0.1:48448", nil)
}

func OutputJson(w http.ResponseWriter, obj interface{}) {
	js, _ := json.Marshal(obj)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
