package main

import "fmt"

var counter int64
var swapcount int64

//////////////////////////////////
//	Erlend Wiklem 2017	//
//	ICA02, IS-105		//
//////////////////////////////////

func main() {
	fmt.Println("Sort this list for me, Erlend!")
	list := []int{6, 14, 4, 18, 17, 1, 20, 7, 3, 15, 5, 2, 16, 10, 19, 8, 11, 13, 9, 12}
	fmt.Println(list)
	BubbleBigO(list)
	Bubble(list)
	fmt.Println("Ok, your list is sorted!")
	fmt.Println(list)
	fmt.Println("")
	fmt.Println(
		"Bubblesort sammenligner første element med andre element, og swapper de hvis første element er større enn andre\n" +
		"Slik fortsetter den loopen til hele listen er sortert\n" +
		"Bubblesort har to looper og har en eksponentiell vekstrate som gir WORST CASE / BIG O = n^2")
}

// Own bubble sort function
func Bubble(list []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(list) -1; i++{
			if list[i + 1] < list[i]{
				DoSwap(list, i, i + 1)
				swap = true
			}
		}
	}
}

// External swap function for re-use
func DoSwap(list []int, i1 ,i2 int){
	tmp := list[i1];
	list[i1] = list[i2];
	list[i2] = tmp;
}

// Find BigO
func BubbleBigO(list []int){
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(list) -1; i++{
			counter++
			if list[i + 1] < list[i]{
				DoSwap(list, i, i + 1)
				swapcount++
				swap = true
			}
		}
	}
	fmt.Println("Total checks:")
	fmt.Println(counter)
	fmt.Println("Total swaps:")
	fmt.Println(swapcount)
}