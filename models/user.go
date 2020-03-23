package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id			int64		`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Phone		string		`json:"phone" orm:"column(phone);unique;null;size(11);description(手机号码)"`
	Username	string		`json:"username" orm:"column(username);unique;size(40);description(用户名称)"`
	Password	string		`json:"-" orm:"column(password);size(40);description(用户密码)"`
	Email		string		`json:"email" orm:"column(email);size(255);description(用户邮箱)"`
	Avatar		string		`json:"avatar" orm:"column(avatar);size(255);null;description(用户头像)"`
	IsActive	string		`json:"is_active" orm:"column(is_active);size(5);default(1);description(是否激活)"`
	Createtime	time.Time	`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
	Updatetime	time.Time	`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
	Department	*Department	`json:"department" orm:"column(department_id);rel(fk);null;description(所属部门)"`
}

// 获取表名，调用的base中的方法，带表名前缀
func (u *User) TableName() string {
	return TableName("user")
}

func init() {
	// 映射model数据（User）
	orm.RegisterModel(new(User))
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// 检测手机号是否注册
func CheckUserPhone(phone string) bool {
	exist := Users().Filter("phone", phone).Exist()
	return exist
}

// 检测用户昵称是否存在
func CheckUserUsername(username string) bool {
	exist := Users().Filter("username", username).Exist()
	return exist
}
// 检测邮箱是否注册
func CheckUserEmail(email string) bool {
	exist := Users().Filter("email", email).Exist()
	return exist
}

//创建用户
func CreateUser(user User) User {
	o := orm.NewOrm()
	o.Insert(&user)
	return user
}

//检测手机和昵称是否注册
func CheckUserPhoneOrUsername(phone string, username string) bool {
	cond := orm.NewCondition()
	count, _ := Users().SetCond(cond.And("phone", phone).Or("username", username)).Count()
	if count <= int64(0) {
		return false
	}
	return true
}
func CheckUserAuth(email string, password string) (User, bool) {
	o := orm.NewOrm()
	user := User{
		Email: email,
		Password: password,
	}
	err := o.Read(&user, "Email", "Password")
	if err != nil {
		return user, false
	}
	return user, true
}

// User database CRUD methods include Insert, Read, Update and Delete
func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}
