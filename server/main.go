package main

import "server/Router"

func main() {
	//user := Model.User{
	//	Username: "tom3",
	//	Password: "123456",
	//}
	//Dao.Mgr.AddUser(&user)
	//Comment := Model.Comment{
	//	Content: "test",
	//	Status:  200,
	//}
	//Dao.Mgr.AddComment(&Comment)
	Router.Start()
}
