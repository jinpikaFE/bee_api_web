package models

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/validation"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	UserList["user_11111"] = &u
}

type User struct {
	Id       string
	Username string `valid:"Required;Match(/^Bee.*/)"`
	Password string `valid:"Required"`
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

// // 如果你的 struct 实现了接口 validation.ValidFormer
// // 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
// func (u *user) Valid(v *validation.Validation) {

// }

func AddUser(u User) interface{} {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	// 验证方法报错
	if err != nil {
		log.Panicln(err)
	}
	// 验证不通过
	if !b {
		errMap := make(map[string]interface{})
		for _, err := range valid.Errors {
			errMap[err.Key] = err.Message
		}
		return &errMap
	}
	return &u
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
