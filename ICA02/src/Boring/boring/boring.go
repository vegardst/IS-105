package boring

import (
	"fmt"
	"time"
	"math/rand"
)



// La til stopp på 100.
func Boring01(msg string) {
	for i := 0; i < 100 ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
func Boring10(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}