package model
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	//"simple_wechat/common/message"
)

//我们在服务器启动后，就初始化一个全局的userDao实例
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool //这里的pool私有
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userdao *UserDao) {
	userdao = &UserDao {
		pool : pool,
	}
	return
}

func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}

	user = &User{}
	//将res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao的连接池中取出一个conn
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userId)
	if err != nil {
		return
	}
	//这时证明这个用户是获取到
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *User) (err error) {
	//先从UserDao的连接池中取出一个conn
	conn := this.pool.Get()
	fmt.Println("Register_1")
	defer conn.Close()
	_, err = this.GetUserById(conn, user.UserId)
	if err == nil { //说明已经有了这个用户id了
		err = ERROR_USER_EXISTS
		return
	} else {
		fmt.Println("GetUserById err=", err)
	}
	fmt.Println("Register_2")
	//这时用户id还没有，完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	fmt.Println("Register_3")
	//入库
	//fmt.Printf("Register_3_user.UserId=%v, data=%v\n", user.UserId, string(data))
	fmt.Println(string(data))
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册信息出错，err=", err)
		return
	}
	fmt.Println("Register_4")
	return
}