package process
import (
	"fmt"
	"go_chatroom/simple_wechat/common/model"
	"go_chatroom/simple_wechat/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*model.User = make(map[int]*model.User)
var curUser model.CurUser //登录成功后，进行初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes * message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &model.User {
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}