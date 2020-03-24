package controllers

import (
	"devops/models"
	"github.com/astaxie/beego/validation"
)
var (
	ErrConfigGitIsExist  = ErrResponse{422001, "配置仓库名称已经存在"}
	ErrConfigProcessIsExist  = ErrResponse{422001, "配置进程名称已经存在"}
)

type ConfigGitController struct {
	BaseController
}
type ConfigProcessController struct {
	BaseController
}
type ConfigController struct {
	BaseController
}

// @Title 新建配置仓库
// @Description 新建配置仓库
// @Param name formData string true "配置仓库名称"
// @Param service_module formData string true "服务模块"
// @Success 200 {object} models.ConfigGit
// @Failure 403 参数错误:缺失或格式错误
// @Faulure 422 名称已经存在
// @router /add [post]
func (this *ConfigGitController) AddConfigGit() {
	name := this.GetString("name")
	service_module := this.GetString("service_module")
	valid := validation.Validation{}
	//表单验证
	valid.Required(name, "name").Message("配置仓库名称必填")
	valid.Required(service_module, "service_module").Message("服务模块必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	if models.CheckConfigGitName(name) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrConfigGitIsExist
		this.ServeJSON()
		return
	}

	config_git := models.ConfigGit{
		Name: name,
	}
	this.Data["json"] = Response{0, "success.", models.CreateConfigGit(config_git)}
	this.ServeJSON()
}

// @Title 新建配置进程
// @Description 新建配置进程
// @Param name formData string true "配置进程名称"
// @Param config_git_id formData int true "配置仓库ID"
// @Success 200 {object} models.ConfigProcess
// @Failure 403 参数错误:缺失或格式错误
// @Faulure 422 名称已经存在
// @router /add [post]
func (this *ConfigProcessController) AddConfigProcess() {
	name := this.GetString("name")
	config_git_id, err := this.GetInt64("config_git_id", -1)
	if err != nil{
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = ErrResponse{403001, map[string]string{"参数错误": "config_git_id未传"}}
		this.ServeJSON()
		return
	}

	valid := validation.Validation{}
	//表单验证
	valid.Required(name, "name").Message("配置进程名称必填")
	valid.Required(config_git_id, "config_git_id").Message("配置仓库ID必传")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	if models.CheckConfigProcessName(name) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrConfigProcessIsExist
		this.ServeJSON()
		return
	}
	config_git := models.FindConfigGitByID(config_git_id)

	config_process := models.ConfigProcess{
		Name: name,
	}
	// 取地址，外键赋值
	config_process.ConfigGit = &config_git
	this.Data["json"] = Response{0, "success.", models.CreateConfigProcess(config_process)}
	this.ServeJSON()
}

// @Title 新建配置
// @Description 新建配置
// @Param config_git_id formData int true "配置仓库ID"
// @Param config_process_id formData int true "配置进程ID"
// @Param content formData json true "配置json内容"
// @Success 200 {object} models.Config
// @Failure 403 参数错误:缺失或格式错误
// @router /add [post]
func (this *ConfigController) AddConfig() {
	content := this.GetString("content")
	config_git_id, _ := this.GetInt64("config_git_id", -1)
	config_process_id, _ := this.GetInt64("config_process_id", -1)

	valid := validation.Validation{}
	//表单验证
	valid.Required(config_git_id, "config_git_id").Message("配置仓库ID必填")
	valid.Required(config_process_id, "config_process_id").Message("配置进程ID必填")
	valid.Required(content, "content").Message("配置内容必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	config_git := models.FindConfigGitByID(config_git_id)
	config_process := models.FindConfigProcessByID(config_process_id)

	config := models.Config{
		Content: content,
	}
	config.ConfigGit = &config_git
	config.ConfigProcess = &config_process

	this.Data["json"] = Response{0, "success.", models.CreateConfig(config)}
	this.ServeJSON()
}