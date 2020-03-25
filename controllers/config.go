package controllers

import (
	"devops/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/globalsign/mgo/bson"
)

const (
	db = "configs"
	collection = "configs_model"
)

type ConfigsController struct {
	BaseController
}

// @Title 查询服务模块
// @Description 查询服务模块
// @Success 200 {object} Response
// @Faulure 422 不存在
// @router /sm [get]
func (this *ConfigsController) ListServiceModule() {
	var service_module interface{}
	if err := models.GetServiceModule(&service_module);err != nil {
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", &service_module}
	this.ServeJSON()
}

// @Title 新建配置
// @Description 新建配置
// @Param service_module formData string true "服务模块"
// @Param content formData json true "配置json内容"
// @Success 200 {object} Response
// @Failure 403 参数错误:缺失或格式错误
// @Faulure 422 名称已经存在
// @router /add [post]
func (this *ConfigsController) AddConfigs() {
	service_module := this.GetString("service_module")
	// 获取传入的json内容
	content := this.GetString("content")
	//fmt.Println(content)
	// json字符串转 golang insterface{}
	var bdoc interface{}
	//err := bson.UnmarshalJSON([]byte(`{"id": 1,"name": "Chummy","age": 19,"tags": ["red", "green"]}`),&bdoc)
	err := bson.UnmarshalJSON([]byte(content),&bdoc)
	bdoc.(map[string]interface{})["service_module"] = service_module
	//fmt.Println(bdoc)
	if err != nil {
		panic(err)
	}
	// 写入mongo
	err = models.Insert(&bdoc)
	if err != nil {
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", &bdoc}
	this.ServeJSON()
}


// @Title 查询Project
// @Description 查询Project
// @Param service_module formData string true "服务模块"
// @Param project_name formData string false "项目名称"
// @Success 200 {object} Response
// @Faulure 422 不存在
// @router /project [get]
func (this *ConfigsController) GetProjects() {
	service_module := this.GetString("service_module")
	project_name := this.GetString("project_name")
	valid := validation.Validation{}
	//表单验证
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
	// 构造查询参数
	var query []bson.M
	if service_module != ""{
		q1 := bson.M{"service_module":service_module}
		query = append(query, q1)
	}
	if project_name != ""{
		q2 := bson.M{"project_name":project_name}
		query = append(query, q2)
	}
	// 构造显示参数(未使用)
	//selector := bson.M{}
	result, err := models.FindProjects(bson.M{"$and": query})
	fmt.Println(result)
	if err != nil{
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", result}
	this.ServeJSON()
}


// @Title 查询进程
// @Description 查询进程
// @Success 200 {object} Response
// @Faulure 422 不存在
// @router /process [get]
func (this *ConfigsController) GetConfigs() {
	result, err := models.FindProcess()
	if err != nil{
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", result}
	this.ServeJSON()
}