package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func httpGet(){
	resp, _ := http.Get("http://localhost:5000/login?username=admin&password=111111")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:5000/login",
		url.Values{"username": {"admin"}, "password": {"111111"}})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

