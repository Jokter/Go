package main

import (
	"goProject/go/src/test2/controller"
	"net/http"
)

/*
a.URI : /mysql/insert
b.参数 id，name，text
c.功能：根据输入参数新增一条mysql记录
 */
/*
a.URI : /mysql/select
b.参数 id
c.功能：根据输入参数 id 查询MySQL数据并返回
*/
func main()  {

	http.HandleFunc("/mysql/select",controller.GetUser)
	http.HandleFunc("/mysql/insert",controller.AddUser)

	//创建路由
	http.ListenAndServe(":8803",nil)
}