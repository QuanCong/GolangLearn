package main


/**
数组切片
 */

func main() {

	var int_array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int = int_array[:]//基于array的所有元素生成切片
	println("切片的长度:", len(slice))

	slice = int_array[:5]
	println("切片的长度:", len(slice))//基于array的前五个元素生成切片

	slice = int_array[1:]//基于array=1以后的所有元素生成切片
	println("切片的长度:", len(slice))

	var slice_make []int = make([]int, 5)//初始元素为5个，元素初始化值为0
	travelSlice(slice_make);
	slice_make2 :=make([]int,5,10);//初始元素为5个，预留10个长度，元素初始值位0
	travelSlice(slice_make2);
	slice_make3 :=[]int{1,2,3,4,5}//直接初始化包含五个元素的切片
	travelSlice(slice_make3);

}

func travelSlice(slice []int) {
	// 传统遍历法
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	//range遍历法
	for i,v:= range  slice{
		println("第",i,"个元素:",v)
	}
}
