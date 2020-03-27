package models

import "github.com/astaxie/beego/orm"


// Menu 菜单资源 实体类
type Menu struct {
	Id			int64		`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Name		string		`json:"name" orm:"column(name);unique;size(40);description(菜单名称)"`
	Rtype		int			`json:"rtype" orm:"column(rtype);description(菜单类型)"`
	Seq			int			`json:"seq" orm:"column(seq);description(菜单排序)"`
	Icon        string		`json:"icon" orm:"column(icon);size(40);description(菜单图标)"`
	LinkUrl     string		`json:"linkurl" orm:"column(linkurl);description(菜单链接)"`
	Parent		*Menu		`json:"parent_id" orm:"rel(fk);default(nil);description(父级菜单)"`//自关联
	Role		[]*Role		`orm:"reverse(many)"` //反向一对多关联
}

// 获取表名，调用的base中的方法，带表名前缀
func (u *Menu) TableName() string {
	return TableName("menu")
}

func init() {
	// 映射model数据（User）
	orm.RegisterModel(new(Menu))
}

func Menus() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Menu))
}