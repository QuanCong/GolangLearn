package website

import "fmt"

type Counter int

func (countP *Counter) Increase() {
	*countP++
}
func (countP *Counter) Decrease() {
	*countP--
}

func (count Counter) IsZero() bool {
	return count == 0
}

type Website struct {
	url  string //具名字段 聚合
	size int    //具名字段  聚合
	Counter     //匿名字段  嵌入
}
//以包的导出方式设定构造方法
func New(url string, size int, count Counter) *Website {
	return &Website{url, size, count}
}
func (site *Website)GetUrl()string  {
	return site.url
}
func (site *Website)SetUrl(url string ){
	site.url=url;
}
func (site Website)IsZero()bool{
	fmt.Println("重写方法")
	return site.Counter==0;
}
