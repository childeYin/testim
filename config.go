package main

type Vertex struct {
	Lat, Long float64
}

type User struct{
	nickName string
	email string
	pwd string
}
var userInfo map[string]User

func init(){
	userInfo = make(map[string]User)
	userInfo["zhangjun"] = User{"尹少爷", "zhangjun", "123456"}
	userInfo["shaoye"]   = User{"小少爷", "小少爷", "123456"}
}