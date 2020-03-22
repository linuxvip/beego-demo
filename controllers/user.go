package controllers

import (
	"devops/models"
	"devops/utils"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

var (
	ErrPhoneIsRegis     = ErrResponse{422001, "手机用户已经注册"}
	ErrUsernameIsRegis  = ErrResponse{422002, "用户名已经被注册"}
	ErrEmailIsRegis  = ErrResponse{422002, "邮箱地址已经被注册"}
	ErrEmailOrPasswd = ErrResponse{422003, "邮箱地址或密码错误。"}
)

type UserController struct {
	BaseController
}
type LoginToken struct {
	Email  models.User `json:"email"`
	Token string      `json:"token"`
}

// @Title 注册新用户
// @Description 用户注册
// @Param username formData string true "用户名称"
// @Param password formData string true "密码"
// @Param email formData string	true "邮箱"
// @Param phone formData string	true "用户手机号"
// @Success 200 {object} models.User
// @Failure 403 参数错误:缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (this *UserController) Registered() {
	//var a models.User
	//json.Unmarshal(this.Ctx.Input.RequestBody, &a)
	//username := a.Username
	//password := a.Password
	//email := a.Email
	//phone := a.Phone

	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	phone := this.GetString("phone")

	valid := validation.Validation{}
	//表单验证
	valid.Required(username, "username").Message("用户名必填")
	valid.Required(password, "password").Message("密码必填")
	valid.Required(phone, "phone").Message("手机必填")
	valid.Required(email, "email").Message("邮箱必填")
	valid.Mobile(phone, "phone").Message("手机号码不正确")
	valid.Email(email, "email").Message("邮箱格式不正确")
	valid.MinSize(username, 2, "username").Message("用户名最小长度为 2")
	valid.MaxSize(username, 40, "username").Message("用户名最大长度为 40")
	valid.MinSize(password, 8, "password").Message("密码最小长度为 8")
	valid.MaxSize(password, 40, "password").Message("密码最大长度为 40")
	//valid.Length(password, 8, "password").Message("密码格式不对")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	if models.CheckUserPhone(phone) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrPhoneIsRegis
		this.ServeJSON()
		return
	}
	if models.CheckUserUsername(username) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrUsernameIsRegis
		this.ServeJSON()
		return
	}
	if models.CheckUserEmail(email) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrEmailIsRegis
		this.ServeJSON()
		return
	}
	// 密码加密
	password = utils.Secret2Password(email, password)

	user := models.User{
		Phone:    phone,
		Username: username,
		Password: password,
		Email: email,
		// 默认为激活状态（0 = 激活，1=冻结）
		IsActive: "0",
	}
	this.Data["json"] = Response{0, "success.", models.CreateUser(user)}
	this.ServeJSON()
}

// @Title 登录
// @Description 账号登录
// @Success 200 {object} models.User
// @Failure 404 no enough input
// @Failure 401 No Admin
// @router /login [post]
func (this *UserController) Login() {
	email := this.GetString("email")
	password := this.GetString("password")

	valid := validation.Validation{}
	//表单验证
	valid.Required(email, "email").Message("邮箱必填")
	valid.Required(password, "password").Message("密码必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}

	// 密码加密
	password = utils.Secret2Password(email, password)
	// 使用邮箱和密码去验证
	user, ok := models.CheckUserAuth(email, password)
	if !ok {
		this.Data["json"] = ErrEmailOrPasswd
		this.ServeJSON()
		return
	}

	et := utils.EasyToken{
		Email: user.Email,
		Uid:      user.Id,
		Expires:  time.Now().Unix() + 3600,
	}

	token, err := et.GetToken()
	if token == "" || err != nil {
		this.Data["json"] = ErrResponse{-0, err}
	} else {
		this.Data["json"] = Response{0, "success.", LoginToken{user, token}}
	}

	this.ServeJSON()
}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {object} models.User
// @Failure 401 unauthorized
// @router /auth [get]
func (this *UserController) Auth() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(authtoken)
	if !valido {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = ErrResponse{-1, fmt.Sprintf("%s", err)}
		this.ServeJSON()
		return
	}

	this.Data["json"] = Response{0, "success.", "user is login"}
	this.ServeJSON()
}


// @Title 头像上传
// @Description 头像上传
// @Param	avatar  query  []byte	false   获取图片二进制流出
// @Success 200
// @router /upload_avatar [post]
func (this *UserController) UploadAvatar() {
	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", this.Ctx.Request.Header.Get("Origin"))

	tmpfile, fheader, err  := this.Ctx.Request.FormFile("avatar")
	// u.GetFile("avatar") 效果相同  “avatar”是二进制流的键名.获取上传的文件
	if err != nil{
		panic(err)
	}
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer tmpfile.Close()
	path := "20181212.jpg"  //设置保存路径
	fmt.Println(fheader.Header)
	fmt.Println(fheader.Size)
	fmt.Println(fheader.Filename)
	//beego.Info("Header:", fheader.Header)//map[Content-Disposition:[form-data; name="123"; filename="upimage.jpg"] Content-Type:[image/jpeg]]
	//beego.Info("Size:", fheader.Size)    //114353
	//beego.Info("Filename:", fheader.Filename)  //upimage.jpg
	this.SaveToFile("123", path)
}