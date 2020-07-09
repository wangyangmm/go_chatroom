package process

import (
	"fmt"
	"encoding/json"
	"go_chatroom/simple_wechat/common/message"
	"go_chatroom/simple_wechat/common/utils"
)

type ClientSmsProcess struct {

}

func (this *ClientSmsProcess) SendGroupMes(content string) (err error) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsGroupMesType

	//2.创建一个SmsMes 实例
	var smsGroupMes message.SmsGroupMes
	smsGroupMes.Content = content
	smsGroupMes.UserId = curUser.UserId
	smsGroupMes.UserStatus = curUser.UserStatus

	//3.序列化 smsGroupMes
	data, err := json.Marshal(smsGroupMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	//5.将mes发送给服务器
	tf := &utils.Transfer {
		Conn : curUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg err=", err)
		return
	}
	return
}

func (this *ClientSmsProcess) SendMesToEachOnlineUser(content string, recvUserId int) {
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.RecvUserId = recvUserId
	smsMes.SendUserId = curUser.UserId

	//3.序列化 smsGroupMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser json.Marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	//4.对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser json.Marshal fail =", err.Error())
		return
	}

	//5.将mes发送给服务器
	tf := &utils.Transfer {
		Conn : curUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser WritePkg err=", err)
		return
	}
	return
}