package main

import "fmt"

var pc [8]byte

func main() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	fmt.Println(pc)
	var sum int

	for i := 0; i < len(pc); i++ {
		sum += PopCount(uint64(pc[i]))
	}
	fmt.Println(sum)
}

func PopCount(x uint64) int {
	var sum byte
	for i := 0; i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}