package main

import ()
import (
	"os"
	"fmt"
)

func main() {

	for i, arg := range os.Args {
		fmt.Println("第",i,"个参数:",arg)
	}

}
