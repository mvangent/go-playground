package main

import "fmt"

func timesTwo(number int) (int, int) {
	return number, number * 2
}

func main() {
	n, nn := timesTwo(4)
	fmt.Println(n, nn)
}
