package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T)  {
	fmt.Println("测试user中的相关方法")
	t.Run("",testAddUser)
	t.Run("",testGetUserById)
}

func testAddUser(t *testing.T){
	AddUser("wangyu","123456")

}
//
func testGetUserById(t *testing.T){
	user,_ := GetUserById(2)
	fmt.Println("获取的用户信息是：",user)
}
//
//func testGetUsers(t *testing.T){
//	fmt.Println("测试查询所有记录")
//	user := &User{}
//	users,_ := user.GetUsers()
//
//	for k,v := range users{
//		fmt.Printf("第%v个用户是%v\n",k,v)
//	}
//}