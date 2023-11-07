package echo2

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// 直到拿不到值退出循环
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
