package packagesinit

import "fmt"
var Value int

func init() {

	fmt.Println("packages init...")
	Value = 10;
}
