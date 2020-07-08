package model
import (
	"net"
)

//在客户端很多地方会用到curUser，做成全局的
type CurUser struct {
	Conn net.Conn
	User
}