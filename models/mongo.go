package models

import (
	"github.com/astaxie/beego/config"
	"github.com/globalsign/mgo"
	"log"
)

//type Configs struct {
//	Id			int64		`json:"id" orm:"column(id);pk;auto;unique;description(id)"`
//	Name		string		`json:"name" orm:"column(name);unique;null;size(11);description(名称)"`
//	Createtime	time.Time	`json:"create_time" orm:"column(create_time);auto_now_add;type(datetime);description(创建时间)"`
//	Updatetime	time.Time	`json:"update_time" orm:"column(update_time);auto_now;type(datetime);description(更新时间)"`
//}


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
	}else{
		log.Println("MongoDB create session succeed")
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

// 配置插入
func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}