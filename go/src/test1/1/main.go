package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello world")
}

func main()  {
	http.HandleFunc("/",handler)

	//创建路由
	http.ListenAndServe(":8803",nil)
}

/*
package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"Hello world")
}

func main()  {
	myHandler := MyHandler{}
	http.Handle("/myHandler",&myHandler)
	http.ListenAndServe(":8803",nil)
}
 */