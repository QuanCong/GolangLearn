package main

import (
	"os"
	"fmt"
	"os/exec"
	"path"
)

/**
获取文件路径
 */
func main() {

	file1,_:=os.Getwd()
	fmt.Println(file1)

	file2,_:=exec.LookPath(os.Args[0])
	fmt.Println(file2)

	dir,_:=path.Split(file2)
	fmt.Println(dir)

	os.Chdir(dir)

	wd,_:=os.Getwd()
	fmt.Println(wd)

}
