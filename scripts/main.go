package main

import (
    utils "devops/scripts/uitls"
    "encoding/json"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
    "path"
)

func respHandler(res interface{}) (tmp map[string]interface{}) {
    // map 需要初始化一个出来
    tmp = make(map[string]interface{})
    log.Println("input res is : ", res)
    switch res.(type) {
    case nil:
        return tmp
    case map[string]interface{}:
        return res.(map[string]interface{})
    case map[interface{}]interface{}:
        log.Println("map[interface{}]interface{} res:", res)
        for k, v := range res.(map[interface{}]interface{}) {
            log.Println("loop:", k, v)
            switch k.(type) {
            case string:
                switch v.(type) {
                case map[interface{}]interface{}:
                    log.Println("map[interface{}]interface{} v:", v)
                    tmp[k.(string)] = respHandler(v)
                    continue
                default:
                    log.Printf("default v: %v %v \n", k, v)
                    tmp[k.(string)] = v
                }

            default:
                continue
            }
        }
        return tmp
    default:
        // 暂时没遇到更复杂的数据
        log.Println("unknow data:", res)
    }
    return tmp
}

func main() {
    // 获取yaml配置目录
    dir, _ := os.Getwd()
    yaml_dir := path.Join(dir, "scripts/yaml_dir")

    // 获取目录下配置文件
    list := utils.GetAllFile(yaml_dir)

    // 遍历每个配置，来进行导入操作
    for _, config_file := range list{
        fmt.Println(config_file)
        file, _ := ioutil.ReadFile(config_file)

        var bdoc map[string]interface{}
        if err := yaml.Unmarshal(file, &bdoc); err != nil{
            panic(err)
        }
        for k, v := range bdoc{
            fmt.Printf("============== %s ==============\n", k)
            fmt.Printf("%T\n",v)
            tmp := respHandler(v)
            log.Println("tmp:", tmp)

            by, err := json.Marshal(tmp)
            log.Println("output json:", string(by), err)
        }
    }
}