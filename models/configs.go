package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// config git 表设计
type ConfigGit struct {
	Id			int64		`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Name		string		`json:"name" orm:"column(name);unique;size(40);description(配置仓库名称)"`
	Createtime	time.Time	`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
	Updatetime	time.Time	`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
	ServiceModule string	`json:"service_module" orm:"column(service_module);size(40);description(服务模块)"`
	ConfigProcess	[]*ConfigProcess	`orm:"reverse(many)"` //反向一对多关联
	Config		[]*Config		`orm:"reverse(many)"` //反向一对多关联
}
// config process 表设计
type ConfigProcess struct {
	Id			int64			`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Name		string			`json:"name" orm:"column(name);unique;size(40);description(进程名称)"`
	Createtime	time.Time		`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
	Updatetime	time.Time		`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
	ConfigGit	*ConfigGit		`json:"config_git_id" orm:"rel(fk);default(nil);description(所属仓库)"` // 外键
	Config		[]*Config		`orm:"reverse(many)"` //反向一对多关联

}
// config 表设计
type Config struct {
	Id			int64			`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
	Content 	string			`json:"content" orm:"column(content);description(json内容)"`
	Createtime	time.Time		`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
	Updatetime	time.Time		`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
	ConfigGit	*ConfigGit		`json:"config_git_id" orm:"rel(fk);default(nil);description(所属仓库)"` // 外键
	ConfigProcess	*ConfigProcess	`json:"config_process_id" orm:"rel(fk);default(nil);description(所属进程)"` // 外键
}
// 获取表名，调用的base中的方法，带表名前缀
func (u *ConfigGit) TableName() string {
	return TableName("config_git")
}
func (u *ConfigProcess) TableName() string {
	return TableName("config_process")
}
func (u *Config) TableName() string {
	return TableName("config")
}

func init() {
	// 映射model数据（Role）
	orm.RegisterModel(new(ConfigGit), new(ConfigProcess), new(Config))
}

func ConfigGits() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(ConfigGit))
}
// 检测配置仓库名称是否存在
func CheckConfigGitName(name string) bool {
	exist := ConfigGits().Filter("name", name).Exist()
	return exist
}

// 根据id查找config git
func FindConfigGitByID(id int64) ConfigGit {
	o := orm.NewOrm()
	config_git := ConfigGit{Id: id}
	err := o.Read(&config_git)
	if err != nil {
		panic(err)
	}
	return config_git
}

// 新建ConfigGit
func CreateConfigGit(config_git ConfigGit) ConfigGit{
	o := orm.NewOrm() // 创建一个 Ormer
	id, err := o.Insert(&config_git)
	fmt.Printf("ID: %s created.", id)
	if err != nil{
		panic(err)
	}
	return config_git
}


func ConfigProcesses() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(ConfigProcess))
}
// 检测配置进程名称是否存在
func CheckConfigProcessName(name string) bool {
	exist := ConfigProcesses().Filter("name", name).Exist()
	return exist
}

// 根据id查找config process
func FindConfigProcessByID(id int64) ConfigProcess {
	o := orm.NewOrm()
	config_process := ConfigProcess{Id: id}
	err := o.Read(&config_process)
	if err != nil {
		panic(err)
	}
	return config_process
}

// 新建ConfigProcess
func CreateConfigProcess(config_process ConfigProcess) ConfigProcess{
	o := orm.NewOrm() // 创建一个 Ormer
	id, err := o.Insert(&config_process)
	fmt.Printf("ID: %s created.", id)
	if err != nil{
		panic(err)
	}
	return config_process
}

func Configs() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Config))
}

// 新建Config
func CreateConfig(config Config) Config{
	o := orm.NewOrm() // 创建一个 Ormer
	id, err := o.Insert(&config)
	fmt.Printf("ID: %s created.", id)
	if err != nil{
		panic(err)
	}
	return config
}
