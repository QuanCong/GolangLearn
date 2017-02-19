package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
	"os"
)

func login(writer http.ResponseWriter,request *http.Request)  {


	fmt.Println("method:",request.Method)
	if request.Method== "GET" {
		file,_:=os.Getwd();
		path:=file+"/src/webserver/login.gtpl"
		t,err:=template.ParseFiles(path)
		if err!=nil{
			log.Fatal(err)
		}
		t.Execute(writer, nil)
	}else {
		request.ParseForm()
		fmt.Println("username:",request.Form.Get("username"))
		fmt.Println("password",request.Form.Get("password"))
	}

}


func printParams(writer http.ResponseWriter, request *http.Request) {

	return
	request.ParseForm(); //解析参数 必须放在第一句
	fmt.Println(request.Form)
	//fmt.Println("path", request.URL.Path)
	//fmt.Println("scheme", request.URL.Scheme)

	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println(writer,"hello") //返回客户端
}

func main() {
	http.HandleFunc("/",printParams)//设置访问路由
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":8080",nil)//设置监听的端口
	if err!=nil{
		log.Fatal(err)
	}
	
}
