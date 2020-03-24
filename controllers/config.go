package controllers

import (
	"devops/models"
	"github.com/globalsign/mgo/bson"
)

const (
	db = "configs"
	collection = "configs_model"
)

type ConfigsController struct {
	BaseController
}

// @Title 新建配置
// @Description 新建配置
// @Param content formData json true "配置json内容"
// @Success 200 {object} Response
// @Failure 403 参数错误:缺失或格式错误
// @Faulure 422 名称已经存在
// @router /add [post]
func (this *ConfigsController) AddConfigs() {
	// 获取传入的json内容
	content := this.GetString("content")

	// json字符串转 golang insterface{}
	var bdoc interface{}
	//err := bson.UnmarshalJSON([]byte(`{"id": 1,"name": "刘成明","age": 29,"tags": ["home", "green"]}`),&bdoc)
	err := bson.UnmarshalJSON([]byte(content),&bdoc)
	if err != nil {
		panic(err)
	}

	// 写入mongo
	err = models.Insert(db, collection, &bdoc)
	if err != nil {
		panic(err)
	}
	// 返回api
	this.Data["json"] = Response{0, "success.", &bdoc}
	this.ServeJSON()
}