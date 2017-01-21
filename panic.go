package main

import (
	"math"
	"fmt"
)

func main() {
	if value, err := ConverWrapper(8); err != nil {
		fmt.Printf("转换出错：%v", err)
	} else {
		fmt.Println(value)
	}

}

func Conver2Int(value int64) (ret int) {
	if value > math.MinInt32&&value < math.MaxInt32 {
		ret = int(value)
		return ret
	} else {
		panic(fmt.Sprintf("%d is not in int range", value))

	}
}

func ConverWrapper(value int64) (ret int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	ret = Conver2Int(value)
	fmt.Println("顺利执行")
	if e := recover(); e != nil {
		err = fmt.Errorf("%v", e)
	}
	return ret, nil

}
