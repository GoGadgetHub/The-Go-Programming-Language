package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		input.Scan()           // 读取下一行，读取到内容返回true，没有更多内容返回false
		counts[input.Text()]++ // 读取到下一行的内容
	}

	// 键，值 = range map
	for line, n := range counts {
		fmt.Println("line", line)
		fmt.Println("n", n)
	}
}
