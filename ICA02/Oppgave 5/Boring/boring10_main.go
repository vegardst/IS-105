package main

import (
	"./boring"
	"fmt"
)

func main() {
	chanstring := make(chan string)
	go boring.Boring10("Boring!", chanstring)
	for i := 0; i < 100 ; i++ {
		fmt.Printf("Hei boring %q\n", <-chanstring)
	}
	fmt.Println("Ferdig!")
}
