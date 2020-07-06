package process
import (
	"fmt"
	"net"
	"simple_wechat/common/message"
	"simple_wechat/common/utils"
	"simple_wechat/common/model"
	"encoding/json"
	
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}

//通知所有在线用户
func (this *UserProcess) NotifyOthersOnineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {//过滤自己
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//对message再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送，创建Transfer实例
	tf := &utils.Transfer {
		Conn : this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (this *UserProcess)ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从message中取出 mes.Data，并反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err=", err)
		return
	}

	//声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//再声明一个 LoginResMes
	var loginResMes message.LoginResMes
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = err.Error()
		}


	} else {
		loginResMes.Code = 200
		//将登录成功的userid赋值给this的UserId字段
		this.UserId = loginMes.UserId
		//登录成功，将登录成功的用户放到userMgr中
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnineUser(loginMes.UserId)
		//将当前在线用户的id 放入loginResMes.UsersId
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println(user, "登录成功")
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//6. 发送data
	tf := &utils.Transfer {
		Conn : this.Conn,
	}

	err = tf.WritePkg(data)
	return
}

func (this *UserProcess)ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从message中取出 mes.Data，并反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err=", err)
		return
	}

	//声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//再声明一个 LoginResMes
	var registerResMes message.RegisterResMes

	fmt.Println("before register:user=", registerMes.User)
	err = model.MyUserDao.Register(&registerMes.User)
	fmt.Println("after register")

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册，未知错误"
		}


	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.User, "注册成功")
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//6. 发送data
	tf := &utils.Transfer {
		Conn : this.Conn,
	}

	err = tf.WritePkg(data)
	return
}
