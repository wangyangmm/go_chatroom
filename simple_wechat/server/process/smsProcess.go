package process
import (
	"fmt"
	"net"
	"go_chatroom/simple_wechat/common/message"
	"go_chatroom/simple_wechat/common/utils"
	"encoding/json"
)

type ServerSmsProcess struct {
	//...暂时不需要字段，说不定以后会有
}

func (this *ServerSmsProcess) SendMes(mes *message.Message) {
	//取出msg中的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.RecvUserId { //找到目标用户，进行转发消息
			this.SendMesToEachOnlineUser(data, up.Conn)
			break
		}
	}
	return
}

func (this *ServerSmsProcess) SendGroupMes(mes *message.Message) {

	//取出msg中的内容
	var smsGroupMes message.SmsGroupMes
	err := json.Unmarshal([]byte(mes.Data), &smsGroupMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsGroupMes.UserId { //过滤自己
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
	return
}

func (this *ServerSmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//发送，创建Transfer实例
	tf := &utils.Transfer {
		Conn : conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
		return
	}
	return
}