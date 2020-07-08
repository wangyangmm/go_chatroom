package main
import (
	"fmt"
	"net"
	"go_chatroom/simple_wechat/server/process"
	"go_chatroom/simple_wechat/common/model"
	"time"
)

func initUserDao() {
	//这里的pool是redis.go中的全局变量
	//注意初始化顺序，先初始化initPool，再初始化initUserDao
	model.MyUserDao = model.NewUserDao(g_pool)
}

func main() {
	//初始化redis连接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
	
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close() //延时关闭
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端连接服务器....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		} else {
			//连接成功，启动一个协程和客户端保持通讯
			processor := &process.Processor {
				Conn : conn,
			}
			go processor.Process()
		}

	}
}