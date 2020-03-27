package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


// Role 用户角色 实体类
type Role struct {
	Id			int64		`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Name		string		`json:"name" orm:"column(name);unique;size(40);description(角色名称)"`
	Seq			int			`json:"seq" orm:"column(seq);description(菜单排序)"`
	Createtime	time.Time	`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
	Updatetime	time.Time	`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
	Menu		[]*Menu		`json:"menu_id" orm:"column(menu_id);rel(m2m);null;description(菜单)"`
	User		[]*User		`orm:"reverse(many)"` //反向一对多关联
}

// 获取表名，调用的base中的方法，带表名前缀
func (u *Role) TableName() string {
	return TableName("role")
}

func init() {
	// 映射model数据（User）
	orm.RegisterModel(new(Role))
}

func Roles() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Role))
}