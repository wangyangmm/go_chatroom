package process
import (
	"fmt"
	"net"
	"io"
	"go_chatroom/simple_wechat/common/utils"
	"go_chatroom/simple_wechat/common/message"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType :
		//处理登录逻辑
		//创建一个UserProcess实例
		up := &UserProcess {
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType :
		//处理注册逻辑
		//创建一个UserProcess实例
		up := &UserProcess {
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)
		if err != nil {
			fmt.Println("up.ServerProcessRegister err=", err)
			return
		}
	case message.SmsGroupMesType :
		//创建一个SmsProcess实例完成转发群聊消息
		smsProcess := &ServerSmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.SmsMesType :
		//转发私信
		smsProcess := &ServerSmsProcess{}
		smsProcess.SendMes(mes)
	default :
		fmt.Println("消息类型不存在，无法处理！")
	}
	return
}

func (this *Processor) Process() (err error) {
	//延时关闭conn
	defer this.Conn.Close()
	
	//循环读取客户端发来的消息
	for {
		tf := &utils.Transfer {
			Conn : this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		fmt.Println("mes=", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}