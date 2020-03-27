package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

var (
	ErrPhoneIsRegis     = "手机用户已经注册"
	ErrUsernameIsRegis  = "用户名已经被注册"
	ErrEmailIsRegis  = "邮箱地址已经被注册"
	ErrEmailOrPasswd =  "邮箱地址或密码错误"
)

// UserQueryParam 用于查询的类
type UserQueryParam struct {
	BaseQueryParam
	UsernameLike	string //模糊查询
	EmailLike		string //模糊查询
	Phone			string //精确查询
}

type User struct {
	Id			int64		`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Phone		string		`json:"phone" orm:"column(phone);unique;null;size(11);description(手机号码)" valid:"Required; Mobile"`
	Username	string		`json:"username" orm:"column(username);unique;size(40);description(用户名称)" valid:"Required; MinSize(2); MaxSize(40)"`
	Password	string		`json:"password" orm:"column(password);size(40);description(用户密码)" valid:"Required; MinSize(8); MaxSize(40)"`
	Email		string		`json:"email" orm:"column(email);size(255);description(用户邮箱)" valid:"Required; Email; MaxSize(100)"`
	Avatar		string		`json:"avatar" orm:"column(avatar);size(255);null;description(用户头像)"`
	IsActive	string		`json:"is_active" orm:"column(is_active);size(5);default(1);description(是否激活)"`
	Createtime	time.Time	`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
	Updatetime	time.Time	`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
	Department	*Department	`json:"department" orm:"column(department_id);rel(fk);null;description(所属部门)"`
	Role		*Role		`json:"role" orm:"column(role_id);rel(fk);null;description(所属角色)"`
}

// 自定义验证
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *User) Valid(v *validation.Validation) {
	if CheckUserPhone(u.Phone) {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("phone", ErrPhoneIsRegis)
	}
	if CheckUserUsername(u.Username) {
		v.SetError("username", ErrUsernameIsRegis)
	}
	if CheckUserEmail(u.Email) {
		v.SetError("email", ErrEmailIsRegis)
	}
}

// 用户登录验证使用 struct
type LoginParams struct {
	Email		string		`json:"email" valid:"Required; Email"`
	Password	string      `json:"password" valid:"Required; MinSize(8); MaxSize(40)"`
}
// 验证登录完成返回体
type LoginToken struct {
	User		User		`json:"user"`
	Token		string      `json:"token"`
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

// UserPageList 获取分页数据
func UserPageList(params *UserQueryParam) ([]*User, int64) {
	data := make([]*User, 0)
	// 默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	query := Users().Filter("username__icontains", params.UsernameLike)
	query = query.Filter("email__icontains", params.EmailLike)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)

	return data, total
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
// 用户验证
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
