package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,r.URL.Path)
	fmt.Fprintln(w,r.URL.RawQuery)

	//获取请求参数
	r.ParseForm()
	fmt.Fprintln(w,"get表单参数：",r.Form)
}

func main()  {
	http.HandleFunc("/post",handler)

	//创建路由
	http.ListenAndServe(":8803",nil)
}
