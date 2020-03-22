package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 部门表设计
type Department struct {
	Id		int64	`json:"id" orm:"column(id);pk;auto;unique"`
	Name	string	`json:"name" orm:"column(name);unique;size(40);"`
	Desc	string	`json:"desc" orm:"column(desc);size(40)"`
	//Parent	int64	`json:"parent" orm:"column(desc);size(40)"`
	Createtime  time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	Updatetime  time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}

// 获取表名，调用的base中的方法，带表名前缀
func (u *Department) TableName() string {
	return TableName("department")
}

func init() {
	// 映射model数据（Role）
	orm.RegisterModel(new(Department))
}

func Departments() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Department))
}

// 检测部门名称是否存在
func CheckDepartmentName(name string) bool {
	exist := Departments().Filter("name", name).Exist()
	return exist
}

//创建部门
func CreateDepartment(depa Department) Department {
	o := orm.NewOrm()
	// 取址（地址&）
	o.Insert(&depa)
	return depa
}

// 根据id查询单个部门
func FindDepartmentbyId(id int64) *Department {
	if id <= 0 {
		return nil
	}

	depa := Department{Id: id}
	err := orm.NewOrm().Read(&depa, "Id")
	if err != nil {
		return nil
	}
	return &depa
}

//部门列表
func  DepaList()  ([]Department, int64, error)  {
	o  :=  orm.NewOrm()
	var lists []Department
	// 获取 QuerySeter 对象，api_department 为表名
	//num, err :=  o.QueryTable("api_department").All(&lists)

	// 也可以直接使用对象作为表名
	// user := new(Department)
	num, err :=  o.QueryTable(new(Department)).All(&lists)
	return  lists,  num,  err
}

//修改更新
func  DepaUpdate(id int, name string, desc string) ([]Department, int64, error) {
	o  :=  orm.NewOrm()
	var  lists  []Department
	num,  err  :=  o.QueryTable(new(Department)).Filter("Id",  id).Update(orm.Params{
		"name":  name,
		"desc":  desc,
	})
	return  lists,  num,  err
}

//删除操作
func  DepaDel(id  int)  ([]Department, int64, error)  {
	o  :=  orm.NewOrm()
	var lists []Department
	num,  err  :=  o.QueryTable(new(User)).Filter("Id",  id).Delete()
	return lists, num, err
}
