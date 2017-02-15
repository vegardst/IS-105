# Oppgave 4

Programmeringsoppgave:

#### Modifisert Bubble_Sort
```go
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
```

### Benchmark av BubbleSort
![alt text](https://github.com/IS-105-GitGroup/IS-105-Gruppe1/blob/master/ICA02/Oppgave%204/BubbleSort/benchmark.PNG "Wapp")



#### BigO
- Bubblesort sammenligner første element med andre element, og swapper de hvis første element er større enn andre.
- Slik fortsetter den loopen til hele listen er sortert
- Bubblesort har to looper og har en eksponentiell vekstratesom gir WORST CASE / BIG O = n^2

Bevis på BigO får jeg ved å kjøre analyse av tellingene mine i koden:

![alt text](https://github.com/IS-105-GitGroup/IS-105-Gruppe1/blob/master/ICA02/Oppgave%204/BubbleSort/BigO.PNG "Wapp")
