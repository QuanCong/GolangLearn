package main

type person struct{
    name string
    age int
}

func main()  {

	//var person_map map[string] person //声明
	person_map :=make(map[string] person,100)

	person_map=map[string] person{"1":{"abc",23}}


	person_map["1"]=person{"songxu",20}
	person_map["2"]=person{"songxu",20}

	delete(person_map,"1")


	value,ok := person_map["23"]
	if ok{
		println(value.age)
	}


	println(len(person_map))


}
