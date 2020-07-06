package process
import (
	"fmt"
	"os"
	"net"
	"simple_wechat/common/utils"
)

//显示登录成功后的界面
func ShowMenu() {
	for {
		fmt.Println("---------恭喜xxxx登录成功-------")
		fmt.Println("---------1.显示在线用户列表-------")
		fmt.Println("---------2.发送消息-------")
		fmt.Println("---------3.信息列表-------")
		fmt.Println("---------4.退出系统-------")
		fmt.Println("请选择（1-4）：")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("显示用户列表～")
		case 2:
			fmt.Println("发送消息")
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
		fmt.Printf("客户端正在读取服务器发送的消息...")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，进一步处理
		fmt.Printf("mes=%v", mes)
	}
}