package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,r.URL.Path)
	//获取全部请求头数据
	fmt.Println(w,r.Header)

	//获取Content-Type数据
	fmt.Println(w,"Content-Type:",r.Header["Content-Type"])

	//获取请求体
	length := r.ContentLength
	body := make([]byte,length)
	r.Body.Read(body)
	fmt.Fprintln(w,"请求体为：",string(body))

	//获取请求参数
	r.ParseForm()
	fmt.Fprintln(w,"post表单参数：",r.PostForm)
}

func main()  {
	http.HandleFunc("/post",handler)
	//创建路由
	http.ListenAndServe(":8803",nil)
}