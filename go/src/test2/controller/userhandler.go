package controller

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"goProject/go/src/test2/dao"
	"net/http"
	"strconv"
)

//根据id获取数据
func GetUser(w http.ResponseWriter,r *http.Request){
	//创建链接
	conn, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialDatabase(1),
		redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("Connect to redis failed ,cause by >>>", err)
		return
	}
	defer conn.Close()
	//获取int型id
	query := r.URL.Query()
	id := query.Get("id")
	userID,_:= strconv.Atoi(id)

	//查询redis中是否有id对应数据
	exists, err := redis.Bool(conn.Do("EXISTS", userID))
	if err != nil {
		user ,_ :=dao.GetUserById(userID)
		b,_ := json.Marshal(user)
		_, err = conn.Do("SET", userID, string(b))
		if err != nil {
			fmt.Println("redis set value failed >>>", err)
		}
	}
	fmt.Fprintln(w,exists)
}

//插入数据
func AddUser(w http.ResponseWriter,r *http.Request){
	query := r.URL.Query()
	name := query.Get("name")
	text := query.Get("text")
	dao.AddUser(name,text)
}
