package message
import (
	"go_chatroom/simple_wechat/common/model"
)

const (
	LoginMesType			= "LoginMes"
	LoginResMesType 		= "LoginResMes"
	RegisterMesType			= "RegisterMes"
	RegisterResMesType  	= "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType				= "SmsMes"
)

//用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string  `json:"data"`//消息数据
}

type LoginMes struct {
	UserId int `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

type LoginResMes struct {
	Code int  `json:"code"` //状态码
	UsersId []int `json:"usersId"` //保存用户id的切片
	Error string `json:"error"` //错误信息
}

type RegisterMes struct {
	User model.User `json:"user"` //包含一个User类型对象
}

type RegisterResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

//增加一个SmsMes //发送消息
type SmsMes struct {
	Content string `json:"content"`
	model.User //匿名结构体  继承
}