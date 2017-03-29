package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"time"
	"crypto/md5"
)

func getMethod() {
	response, err := http.Get("http://www.taobao.com")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	file, _ := os.Create("D:/taobao.html")
	defer file.Close()

	io.Copy(file, response.Body)
}

func main() {
	var topic="songxutopic"
	var url="http://publictest-rest.ons.aliyun.com/"
	var date=time.Now().Format("2006-01-02 15:04:05 ")//Mon Jan 2 15:04:05 MST 2006
	fmt.Println(date)
	var ak="FueksRhbjLSjAbZ3"
	var sk="vXpwdfLGBUMtZpJ7iRLlTDWNI7SU7X"
	var pid="PID-songxupd"
	var msgBody="hello aliyun "
	h:=md5.New()
	io.WriteString(h,msgBody)
	//发送内容的md5
	var md5=string(h.Sum(nil))
	//签名字符串
	var newline="\n"
	var sign=topic+newline+pid+newline+md5+newline+date
	req,_:=http.NewRequest("POST",url+"message/?topic="+topic+"&time="+date+"&tag=http"+"&key=http", nil)

	req.Header.Add("Signature",sign)
	req.Header.Add("AccessKey",ak)
	req.Header.Add("ProducerID",pid)



	respose,err:=http.Post()



}
