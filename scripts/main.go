package main

import (
    utils "devops/scripts/uitls"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "path"
    "strings"
)

type Config struct {
    Info string
}

func main() {
    // 获取yaml配置目录
    dir, _ := os.Getwd()
    yaml_dir := path.Join(dir, "scripts/yaml_dir")

    // 获取目录下配置文件
    list := utils.GetAllFile(yaml_dir)

    // 遍历每个配置，来进行导入操作
    for _, config_file := range list{
        file_name := strings.Split(config_file, "/")
        // 获取servie_module (ai backend frontend)
        service_module := strings.Split(file_name[len(file_name)-1], ".")[0]
        fmt.Println(service_module)

        // yaml转map
        c := Config{}

        file, _ := ioutil.ReadFile(config_file)

        yaml.Unmarshal(file, &c)
        fmt.Println(c)

        //err = yaml.Unmarshal(file, &config)
        //if err != nil {
        //    return config, err
        //}


    }

}