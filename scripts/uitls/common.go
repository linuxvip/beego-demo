package utils

import (
	"io/ioutil"
)


func GetAllFile(pathname string) []string {
	// make a slice of length 0
	list := make([]string, 0)
	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+"/"+fi.Name())
			GetAllFile(pathname + fi.Name() + "/")
		} else {
			//fmt.Println(fi.Name())
			list = append(list, pathname+"/"+fi.Name())
		}
	}
	return list
}