package process
import (
	"fmt"
	"os"
	"encoding/json"
	"net"
	"go_chatroom/simple_wechat/common/message"
	"go_chatroom/simple_wechat/common/utils"
	"go_chatroom/simple_wechat/common/model"
)

type UserProcess struct {

}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close() //一定要记得close

	//2.发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化，此时的data就是要发送的数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7. 发送login相关消息
	tf := &utils.Transfer {
		Conn : conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送登录信息出错")
		return
	}

	//这里还需要处理服务器返回的消息 
	mes, err = tf.ReadPkg() //mes就是LoginResMes
	if err != nil {
		fmt.Println("readpkg(conn) err=", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化curUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline
		//显示当前在线用户列表
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			//要求不显示自己在线
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//初始化客户端用户列表 onineUsers
			user := &model.User {
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//启动一个协程，用来处理服务器端的消息
		go ServerProcessMes(conn)

		//显示登录成功后的菜单
		ShowMenu()

	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

func (this *UserProcess) Register(userId int,
	 userPwd string, userName string) {
	
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close() //一定要记得close

	//2.发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个 registerMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化，此时的data就是要发送的数据
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7. 发送 register 相关消息
	tf := &utils.Transfer {
		Conn : conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册信息出错")
		return
	}

	//这里还需要处理服务器返回的消息 
	mes, err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil {
		fmt.Println("readpkg(conn) err=", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
	}

	return
}