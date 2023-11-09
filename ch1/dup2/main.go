package main

import "fmt"

func main() {
	counts := make(map[string]int)
	countLines(counts)
	fmt.Println(counts)
}

func countLines(counts map[string]int) {
	counts["why"] = 1
}
