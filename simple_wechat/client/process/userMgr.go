package process
import (
	"fmt"
	"simple_wechat/common/model"
)

//客户端要维护的map
var onlineUsers map[int]*User = make(map[int]*User)

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes * message.notifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user := &message.User {
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}