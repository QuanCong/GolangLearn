package main

import (
	"io"
	"encoding/json"
	"fmt"
	"log"
)

const (
	CAT = iota
	DOG
	LION
)

type Pet struct {
	Name  string
	Color string
	Type  int
}

type Person struct {
	name string //非导出的字段不能被编码
	Age  int
	Pet  Pet
}

func toJson(writer io.Writer, p Person) {
	encoder := json.NewEncoder(writer)
	encoder.Encode(p)
}

func main() {
	/*pet := Pet{"dog", "white", DOG}
	p := Person{"sx", 20, pet}

	bstr,_:=json.Marshal("songxu")
	fmt.Println("输出字符串: ",string(bstr))

	mapV:=map[string]string{"1":"1s","2":"2s"}
	bmap,_:=json.Marshal(mapV)
	fmt.Println("输出map: ",string(bmap))

	bs,_:=json.Marshal(p)
	fmt.Println("输出struct: ",string(bs))

	toJson(os.Stdout, p)//输出到命令行 输出到文件

	var p2 Person
	json.Unmarshal(bs,&p2)
	fmt.Println("解析得到的:",p2)

	var m2 map[string]interface{}
	json.Unmarshal(bmap,&m2)
	fmt.Println("解析得到的:",m2)*/

	jsonStr := "{\"name\":\"songxu\",\"age\":20,\"isStudent\":true,\"height\":182,\"family\":[\"father\",\"mother\"]}"

	var unkown interface{}
	err := json.Unmarshal([]byte(jsonStr), &unkown)
	if err != nil {
		log.Fatal(err)
	}
	value, ok := unkown.(map[string]interface{})
	if ok {
		for k, v := range value {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string value:", v2)
			case bool:
				fmt.Println(k, "is bool value:", v2)
			case []interface{}:
				fmt.Println(k, "is array")
				for i, valueArray := range v2 {
					fmt.Println(i, ":", valueArray)
				}

			}
		}
	}

}



