package controllers

import (
	"devops/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"strings"
)

var (
	ErrDepartmentIsExist  = ErrResponse{422001, "部门名称已经存在"}
)

type DepartmentController struct {
	BaseController
}

// @Title 查询部门列表
// @Description 查询部门列表
// @Success 200 {object} models.Department
// @Failure 500 查询错误
// @router / [get]
func (this *DepartmentController) DepartmentList() {
	list, num, err := models.DepaList()
	fmt.Println(num)
	if err != nil{
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Data["json"] = ErrResponse{403001, "查询错误"}
		this.ServeJSON()
	}
	this.Data["json"] = Response{0, "success.", list}
	this.ServeJSON()
}

// @Title 新建部门
// @Description 新建部门
// @Param name formData string true "部门名称"
// @Param desc formData string false "部门描述"
// @Success 200 {object} models.Department
// @Failure 403 参数错误:缺失或格式错误
// @Faulure 422 部门名称已经存在
// @router /add [post]
func (this *DepartmentController) AddDepartment() {
	name := this.GetString("name")
	desc := this.GetString("desc")

	valid := validation.Validation{}
	//表单验证
	valid.Required(name, "name").Message("部门名称必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	if models.CheckDepartmentName(name) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrDepartmentIsExist
		this.ServeJSON()
		return
	}

	depa := models.Department{
		Name: name,
		Desc: desc,
	}
	this.Data["json"] = Response{0, "success.", models.CreateDepartment(depa)}
	this.ServeJSON()
}

// @Title 查询单个部门
// @Description 查询单个部门
// @Param id formData int true "部门id"
// @Success 200 {object} models.Department
// @Failure 404 参数错误:缺失或格式错误
// @router /:id [get]
//func (this *DepartmentController) FindDepartmentById() {
//	id := this.GetString("id", "-1")
//
//	depa := models.FindDepartmentbyId(id)
//	if depa == nil {
//		this.Ctx.WriteString("Department id not exists")
//		return
//	}
//	this.Data["json"] = Response{0, "success.", depa}
//	this.ServeJSON()
//}

// @Title 修改部门
// @Description 修改部门
// @Param id formData int64 true "部门id"
// @Param name formData string true "部门名称"
// @Param desc formData string false "部门描述"
// @Success 200 {object} models.Department
// @Failure 500 参数错误:缺失或格式错误
// @router /:id [post]
func (this *DepartmentController) UpdateDepartment() {
	id, _ := this.GetInt("id")
	fmt.Println(id)
	name  :=  strings.TrimSpace(this.GetString("name"))
	desc  :=  strings.TrimSpace(this.GetString("desc"))
	lists, num, err := models.DepaUpdate(id,  name,  desc)
	fmt.Println(num)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Data["json"] = ErrResponse{403001, "更新错误"}
		this.ServeJSON()
	}
	this.Data["json"] = Response{0, "success.", lists}
	this.ServeJSON()
}


// @Title 删除部门
// @Description 删除部门
// @Param id formData int64 true "部门id"
// @Success 200 {object} models.Department
// @Failure 500 参数错误:缺失或格式错误
// @router /:id [delete]
func (this *DepartmentController) DelDepartment() {
	id, _ := this.GetInt("id")
	fmt.Println(id)
	lists, num, err := models.DepaDel(id)
	fmt.Println(num)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Data["json"] = ErrResponse{403001, "删除错误"}
		this.ServeJSON()
	}
	this.Data["json"] = Response{0, "success.", lists}
	this.ServeJSON()
}