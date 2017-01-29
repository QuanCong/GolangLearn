package main

import (
	"runtime"
	"os"
	"fmt"
	"path/filepath"
	"sync"
	"log"
	"crypto/sha1"
	"io"
)

/**
利用sha1找到给定路径下所有的重复的文件并打印输出
 */

type fileInfo struct {
	path string
	size int64
	sha1 []byte
}
type resultInfo struct {
	size     int64
	filename string
	path     []string
}

const maxgoroutines = 100

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "-help" {
		fmt.Printf("usage: %s <path>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	infochan := make(chan fileInfo, maxgoroutines * 2)
	findDuplicaties(infochan, os.Args[1])
	resultChan := mergeResult(infochan)
	printResult(resultChan)
}
/**
 */
func findDuplicaties(infochan chan fileInfo, dirPath string) {
	waiter := &sync.WaitGroup{}
	var maxSize = int64(1024 * 32)
	//os.ModeType是非文件类型（如文件夹、连接、管道）的位集合，与该类型&操作返回0说明这是一个普通文件类型
	var walkFun = func(path string, info os.FileInfo, err error) error {
		if err == nil&&info.Size() > 0 &&(info.Mode() & os.ModeType == 0) {
			if info.Size() < maxSize || runtime.NumGoroutine() > maxgoroutines {
				//如果文件大不超过32k或者当前goroutine的数量已经超过100 则在当前线程处理文件
				processFile(path, info, infochan, nil)
			} else {
				waiter.Add(1)
				//开启一个新的goroutine处理文件
				go processFile(path, info, infochan, func() {
					waiter.Done()
				})
			}
		}
		return nil//不返回错误
	}

	filepath.Walk(dirPath, walkFun)
	waiter.Wait()//等待所有的文件处理goroutine完成
	close(infochan)//关闭

}
func processFile(filename string, info os.FileInfo, infochan chan fileInfo, done func()) {
	//如果传入了回调方法，则调用
	if done != nil {
		defer done()
	}
	fmt.Println(info.Name())
	file, err := os.Open(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()
	hash := sha1.New()

	if size, err := io.Copy(hash, file); size != info.Size() || err != nil {
		return
	}
	infochan <- fileInfo{path:filename, size:info.Size(), sha1:hash.Sum(nil)}
}

func mergeResult(fileChan chan fileInfo) map[string]*resultInfo {
	resultmap := make(map[string]*resultInfo)
	format := fmt.Sprint("%%016X:%%%dX", sha1.Size * 2)
	for info := range fileChan {
		key := fmt.Sprint(format, info.size, info.sha1)
		value, found := resultmap[key]
		if !found {
			value = &resultInfo{size:info.size}
			resultmap[key] = value
		}
		value.path = append(value.path, info.path)
	}
	return resultmap
}

func printResult(resultmap map[string]*resultInfo) {
	fmt.Println("==================================")
	var found = false
	for key := range resultmap {
		value := resultmap[key]
		if len(value.path) > 1 {
			//有重复文件
			found = true
			fmt.Printf("%d duplicated files (%s bytes)\n", len(value.path), value.size)
			for _, path := range value.path {
				fmt.Printf("\t%s\n", path)
			}
		}
	}
	if !found {
		fmt.Println("no duplicated files found")
	}

}


