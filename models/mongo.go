package models

import (
	"github.com/astaxie/beego/config"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

// 内联嵌套子集
type ProjectResp struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	ServiceModule string        `json:"service_module" bson:"service_module"`
	ProjectName   string        `json:"project_name" bson:"project_name"`
	ProjectInfo   ProcessResp   `json:"project_info" bson:"project_info"`
}
type ProcessResp struct {
	ProcessName 	string 		`json:"process_name" bson:"process_name"`
	ProcessInfo		interface{} `json:"process_info" bson:"process_info"`
}

var globalS *mgo.Session

func init() {

	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	host := appConf.String("mongodb::host")
	source := appConf.String("mongodb::source")
	user := appConf.String("mongodb::user")
	pass := appConf.String("mongodb::pass")

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Source:   source,
		Username: user,
		Password: pass,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("MongoDB create session error ", err)
	} else {
		log.Println("MongoDB create session succeed")
	}
	globalS = s
}

func GetDBandCollection() (db, collection string) {
	db = "configs"
	collection = "configs_model"
	return db, collection
}

func connect() (*mgo.Session, *mgo.Collection) {
	db, collection := GetDBandCollection()
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

// 获取所有的服务模块
func GetServiceModule(result interface{}) error {
	ms, c := connect()
	defer ms.Close()
	err := c.Find(bson.M{}).Distinct("service_module", result)
	return err
}

// 获取所有的project
func FindProjects(query interface{}) (*[]ProjectResp, error) {
	ms, c := connect()
	defer ms.Close()
	tmp := []ProjectResp{}
	//conditions := bson.M{key: bson.RegEx{find, "i"}}//构造模糊查询字段
	//err := c.Find(bson.M{"service_module": "backend"}).All(&tmp)
	err := c.Find(query).All(&tmp)
	//fmt.Println(tmp)
	return &tmp, err
}

// 配置插入
func Insert(docs ...interface{}) error {
	ms, c := connect()
	defer ms.Close()
	return c.Insert(docs...)
}

// 配置修改
func Update(query, update interface{}) error {
	ms, c := connect()
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(query interface{}) error {
	ms, c := connect()
	defer ms.Close()
	return c.Remove(query)
}

//func FindOne(query, selector, result interface{}) error {
//	ms, c := connect()
//	defer ms.Close()
//	return c.Find(query).Select(selector).One(result)
//}
//
//func FindById(id string) (ProjectResp, error) {
//	var result ProjectResp
//	err := FindOne(bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
//	return result, err
//}


