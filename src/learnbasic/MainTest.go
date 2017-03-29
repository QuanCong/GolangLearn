package main

import (
	"website"
	"fmt"
)

func main() {
	site := website.New("abc.com", 20, 0)
	site.Increase();
	website.ExportMethod();
	fmt.Println("站点的url：", site.GetUrl())
	fmt.Println(site.IsZero())
	fmt.Println("站点的访问量：", site.Counter)
}
