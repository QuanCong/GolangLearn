package main

import "fmt"

//接口名字通常以er结尾
type Listener interface {
	Listen()
}

type Person struct {
	name string
	age  int
}
//实现Listener接口
func (person *Person)Listen() {
	fmt.Println(person.name, "is listening")
}

type Cat struct {
	color string
}
//实现Listener接口
func (cat *Cat) Listen() {
	fmt.Println("a", cat.color, "cat listenned sth and run")
}
//接口类型传入
func WorldListen(listeners ...Listener) {
	for _, l := range listeners {
		l.Listen();
	}
}

func main() {
	person := Person{"songxu", 20}
	cat := Cat{"white"}
	WorldListen(&person, &cat)

}
