package process
import (
	"fmt"
	"os"
	"net"
	"strings"
	"go_chatroom/simple_wechat/common/utils"
	"go_chatroom/simple_wechat/common/message"
	"encoding/json"
)

//显示登录成功后的界面
func ShowMenu() {
	for {
		fmt.Printf("---------恭喜%d登录成功-------\n", curUser.UserId)
		fmt.Println("---------1.显示在线用户列表-------")
		fmt.Println("---------2.发送消息-------")
		fmt.Println("---------3.信息列表-------")
		fmt.Println("---------4.退出系统-------")
		fmt.Println("请选择（1-4）：")
		var key int
		var content string
		smsProcess := &ClientSmsProcess{}
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			//fmt.Println("显示用户列表～")
			outputOnlineUser()
		case 2:
			fmt.Println("---------1.群聊------------")
			fmt.Println("---------2.私聊------------")
			fmt.Println("请选择（1-2）：")
			var slct int
			fmt.Scanf("%d\n", &slct)
			switch slct {
			case 1://群聊
				fmt.Println("你想对大家说点什么：")
				fmt.Scanf("%s\n", &content)
				smsProcess.SendGroupMes(content)
			case 2://私聊
				outputOnlineUser()
				var userId int
				for {
					fmt.Println("请选择聊天对象，输入对应的用户id：")
					fmt.Scanf("%d\n", &userId)
					_, res := onlineUsers[userId]
					if !res {//说明输入的id有误
						fmt.Println("输入的id有误！！！")
					} else {
						fmt.Printf("-------------与%d的聊天--------------\n", userId)
						fmt.Println("----------输入quit可退出当前聊天--------------")
						break
					}
				}

				for {
					fmt.Scanf("%s\n", &content)
					if strings.EqualFold(content, "quit") {
						break
					}
					smsProcess.SendMesToEachOnlineUser(content, userId)
				}
			default:
				fmt.Println("输入的选项有误，请重新输入...")
			}

			
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("您已退出系统....")
			os.Exit(0)
		default:
			fmt.Println("输入的选项有误，请重新输入...")
		}
	}
}

func ServerProcessMes(conn net.Conn) {
	//创建一个transer实例，循环读取服务器端来的消息
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		//fmt.Println("客户端正在读取服务器发送的消息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		//如果读取到消息，进一步处理
		switch mes.Type {
		case message.NotifyUserStatusMesType : //有人上线了
			//1.取出.NotifyUserStatusMes
			//2.把这个用户的信息，状态保存到客户map中， map[int]User
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsGroupMesType : //有人群发消息
			outputGroupMes(&mes)
		case message.SmsMesType : //收到私信
			outputMes(&mes)
		default:
			fmt.Println("服务器返回了未知消息类型")
		}
		//fmt.Printf("mes=%v", mes)
	}
}