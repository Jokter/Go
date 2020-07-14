package dao

import (
	"fmt"
	"goProject/go/src/test2/model"
	"goProject/go/src/test2/utils"
)

////添加User方法一
func AddUser(name string,text string) error{
	//sql语句
	sqlStr := "insert into users(name,text) values(?,?)"
	//执行
	_,err := utils.Db.Exec(sqlStr,name,text)
	if err!=nil{
		fmt.Println("执行异常",err)
		return err
	}
	return nil
}

//根据id查询数据
func GetUserById(id int) (*model.User,error){
	sqlStr := "select id,name,text from users where id = ?"
	row := utils.Db.QueryRow(sqlStr,id)

	user := &model.User{}

	row.Scan(&user.ID,&user.Name,&user.Text)

	return user,nil
}
//
////查询多行数据
//func (user *User) GetUsers() ([]*User,error) {
//	sqlStr := "select id,name,text from users"
//	rows,err:= utils.Db.Query(sqlStr)
//	if err!=nil{
//		return nil,err
//	}
//	var users []*User
//
//	for rows.Next(){
//		var id int
//		var name string
//		var text string
//
//		err := rows.Scan(&id,&name,&text)
//
//		if err!=nil{
//			return nil,err
//		}
//
//		u := &User{
//			ID: id,
//			name: name,
//			text: text,
//		}
//
//		users = append(users,u)
//	}
//	return users,nil
//}