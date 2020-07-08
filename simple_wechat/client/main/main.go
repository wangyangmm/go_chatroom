package main
import (
	"fmt"
	"os"
	"go_chatroom/simple_wechat/client/process"
)

//定义两个全局变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int

	for {
		fmt.Println("-------------------------欢迎登录多人聊天系统--------------------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s", &userPwd)
			//完成登录
			//1.创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名:")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}		
	}
}