package process
import (
	"fmt"
)

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = & UserMgr {
		onlineUsers : make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up * UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线用户列表
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	//如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok {//表示查找的用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}