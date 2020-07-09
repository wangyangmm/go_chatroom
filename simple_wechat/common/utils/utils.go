package utils
import (
	"fmt"
	"net"
	"go_chatroom/simple_wechat/common/message"
	"encoding/json"
	"encoding/binary"
)

//将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte  //传输时的缓冲
}

func (this *Transfer)ReadPkg() (mes message.Message, err error) {
	//fmt.Println("读取客户端发送的数据...")
	//读取pkg的header
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	_, err = this.Conn.Read(this.Buf[:4]) 
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//读取pkg的body
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//技术就是一层窗户纸 &mes !!!!
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		//err = errors.New("json unmarshal error")
		return
	}
	return
}

func (this *Transfer)WritePkg(data []byte) (err error) {
	//发送data的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	n, err := this.Conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail:", err)
		return
	}

	//发送data
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return
}