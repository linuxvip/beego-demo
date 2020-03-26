package controllers

import (
	"devops/models"
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

// @Title 新增配置
// @Description 新增配置
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


// @Title 查询配置
// @Description 查询配置
// @Param service_module formData string true "服务模块"
// @Param project_name formData string false "项目名称"
// @Success 200 {object} Response
// @Faulure 422 不存在
// @router /list [get]
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
	//fmt.Println(result)
	if err != nil{
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", result}
	this.ServeJSON()
}


// @Title 修改配置
// @Description 修改配置
// @Param id formData string true "需要更新的ID"
// @Param content formData string true "需要更新的内容(Json)"
// @Success 200 {object} Response
// @Faulure 422 不存在
// @router /update [post]
func (this *ConfigsController) UpdateConfigs() {
	id := this.GetString("id")
	content := this.GetString("content")
	// 通过_id 构造查询，获取需要修改的源数据
	query := bson.M{"_id":bson.ObjectIdHex(id)}

	// 处理提交的 需要更新的数据，json字符串转 golang insterface{}
	var bdoc interface{}
	err := bson.UnmarshalJSON([]byte(content),&bdoc)
	//fmt.Println(bdoc)
	if err != nil {
		panic(err)
	}
	// 调用修改
	err = models.Update(query, &bdoc)
	if err != nil {
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", &bdoc}
	this.ServeJSON()
}


// @Title 修改配置
// @Description 修改配置
// @Param id formData string true "需要删除的ID"
// @Success 200 {object} Response
// @Faulure 422 不存在
// @router /delete [post]
func (this *ConfigsController) DeleteConfigs() {
	id := this.GetString("id")
	// 通过_id 构造查询，获取需要删除的数据
	query := bson.M{"_id":bson.ObjectIdHex(id)}

	// 调用删除
	err := models.Remove(query)
	if err != nil {
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", "删除成功"}
	this.ServeJSON()
}