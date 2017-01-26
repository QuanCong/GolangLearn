package main
// 利用并发实现grep功能
import (
	"runtime"
	"os"
	"fmt"
	"path/filepath"
	"regexp"
	"log"
	"bufio"
	"bytes"
	"io"
)

type Result struct {
	filename string //文件名
	lino     int    //行号
	line     string //匹配文件行
}

type Job struct {
	filename string
	results  chan <- Result
}
//实际完成匹配工作
func (job Job)Do(lineRx *regexp.Regexp) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)

		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := reader.ReadBytes('\n');
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {
			//fmt.Printf("匹配 %s:%d:%s", job.filename, lino, string(line))
			job.results <- Result{job.filename, lino, string(line)}
		}

		if err != nil {
			if err != io.EOF {
				//读取到结尾
				log.Println(err)
			}
			break;

		}

	}

}

func main() {
	fmt.Println("参数：", os.Args)
	runtime.GOMAXPROCS(runtime.NumCPU())//利用所有的CPU核数
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "-help" {
		fmt.Printf("usage:%s <regexp><files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatal("invalid regexp:%s\n", err)
	} else {
		grep(lineRx, os.Args[2:])
	}
}
//正则表达式匹配查找文件
func grep(lineRx *regexp.Regexp, filenames []string) {
	workers := runtime.NumCPU()
	jobs := make(chan Job, workers)
	results := make(chan Result, len(filenames))
	done := make(chan struct{}, workers)
	go addJob(filenames, jobs, results)
	for i := 0; i < workers; i++ {
		go doJobs(jobs, done, lineRx)
	}
	for working := workers; working > 0; {
		select { //阻塞等待所有goroutine完成
		case result := <-results:
			fmt.Printf("select 打印%s:%d:%s\n", result.filename, result.lino, result.line)
		case <-done:
			working--
			fmt.Println("一个工作进程已经完成")
		}
	}
	//打印结果
	Done:
		select {
		case result := <-results:
			fmt.Printf("结果打印：%s:%d:%s\n", result.filename, result.lino, result.line)
		default:
			break Done
		}
}

func addJob(filenames []string, jobs chan <- Job, results chan <-Result) {
	for _, value := range filenames {
		jobs <- Job{value, results}

	}
	close(jobs)//关闭通道来通知接收方所有的job都已经添加完成
}
func doJobs(jobs chan Job, done chan <-  struct{}, lineRex *regexp.Regexp) {
	for job := range jobs {
		job.Do(lineRex)
	}
	done <- struct {

	}{}
}