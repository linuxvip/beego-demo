package controllers

import (
	"devops/models"
	"devops/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
)

var (
	ErrEmailOrPasswd =  "邮箱地址或密码错误"
)


type UserController struct {
	BaseController
}



// @Title 添加用户
// @Description 添加用户
// @Param username formData string true "用户名称"
// @Param password formData string true "密码"
// @Param email formData string	true "邮箱"
// @Param phone formData string	true "用户手机号"
// @Success 200 {object} models.User
// @Failure 423 参数错误:缺失或格式错误
// @router /add [post]
func (this *UserController) AddUser() {
	var u models.User
	valid := validation.Validation{}

	// 获取前端传过来的json的数据
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &u)
	if err != nil {
		panic(err)
	}
	// 根据 User Struct tag 做验证
	HasErr, _ := valid.Valid(&u)

	// 如有错误，直接返回错误响应(model中 还存自定义验证)
	if !HasErr{
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(423)
			this.Data["json"] = ErrResponse{423, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	// 密码加密存储
	u.Password = utils.Secret2Password(u.Email, u.Password)
	user := models.User{
		Phone:u.Phone,
		Username: u.Username,
		Password: u.Password,
		Email: u.Email,
		// 设置默认为激活状态（0 = 激活，1=冻结）
		IsActive: "0",
	}
	this.Data["json"] = Response{0, "success.", models.CreateUser(user)}
	this.ServeJSON()
}

// @Title 用户登录
// @Description 用户登录
// @Param email formData string	true "邮箱地址"
// @Param password formData string true "密码"
// @Success 200 {object} models.LoginToken
// @Failure 423 参数错误:缺失或格式错误
// @Failure 422 验证错误:邮箱或密码错误
// @router /login [post]
func (this *UserController) Login() {
	var u models.LoginParams
	valid := validation.Validation{}

	// 获取前端传过来的json的数据
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &u)
	if err != nil {
		panic(err)
	}
	// 根据 User Struct tag 做验证
	HasErr, _ := valid.Valid(&u)

	// 如有错误，直接返回错误响应(model中 还存自定义验证)
	if !HasErr{
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(423)
			this.Data["json"] = ErrResponse{423, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}

	// 密码加密
	u.Password = utils.Secret2Password(u.Email, u.Password)
	// 使用邮箱和密码去验证
	user, ok := models.CheckUserAuth(u.Email, u.Password)
	if !ok {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrResponse{422, map[string]string{"error": ErrEmailOrPasswd}}
		this.ServeJSON()
		this.StopRun()
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
		this.Data["json"] = Response{0, "success.", models.LoginToken{user, token}}
	}
	this.ServeJSON()
}

// @Title 用户认证
// @Description 用户认证
// @Success 200 {object} Response
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


// @Title 用户列表
// @Description 用户列表
// @Success 200 {object} models.User
// @Failure 401 unauthorized
// @router /list [get]
func (this *UserController) GetUserList()  {
	var params models.UserQueryParam
	params.Limit, _ = this.GetInt64("limit", 10)
	params.Offset, _ = this.GetInt64("offset", 0)
	params.UsernameLike = this.GetString("username")
	params.EmailLike = this.GetString("email")
	//fmt.Println(params)

	data, total := models.UserPageList(&params)
	page_data := PageData{Count:total,Data:data}
	this.Data["json"] = ResponsePage{0, "success.",&page_data}

	this.ServeJSON()
}


//// @Title 头像上传
//// @Description 头像上传
//// @Param	avatar  query  []byte	false   获取图片二进制流出
//// @Success 200
//// @router /upload_avatar [post]
//func (this *UserController) UploadAvatar() {
//	this.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", this.Ctx.Request.Header.Get("Origin"))
//
//	tmpfile, fheader, err  := this.Ctx.Request.FormFile("avatar")
//	// u.GetFile("avatar") 效果相同  “avatar”是二进制流的键名.获取上传的文件
//	if err != nil{
//		panic(err)
//	}
//	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
//	defer tmpfile.Close()
//	path := "20181212.jpg"  //设置保存路径
//	fmt.Println(fheader.Header)
//	fmt.Println(fheader.Size)
//	fmt.Println(fheader.Filename)
//	//beego.Info("Header:", fheader.Header)//map[Content-Disposition:[form-data; name="123"; filename="upimage.jpg"] Content-Type:[image/jpeg]]
//	//beego.Info("Size:", fheader.Size)    //114353
//	//beego.Info("Filename:", fheader.Filename)  //upimage.jpg
//	this.SaveToFile("123", path)
//}