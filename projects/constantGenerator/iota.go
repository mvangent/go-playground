package main

import "fmt"

const (
	hiphop int = 80 + iota
	lindyhop
	contemporary
)

func main() {
	fmt.Printf("hiphop %v\n", hiphop)
	fmt.Printf("lindyhop %v\n", lindyhop)
	fmt.Printf("contemporary %v\n", contemporary)
}
