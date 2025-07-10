package main

import "fmt"

func main() {
	fmt.Println("in Main")
	go myNew()
	myTwo()
}
