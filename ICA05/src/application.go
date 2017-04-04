package main
//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"./restful"
	"fmt"
)

// Simpel applikasjon som bruker samme filer som serveren bruker.
func main() {
	fmt.Print("Skriv lokasjon du vil søke på: ")
	var city string = "Kristiansand"
	fmt.Scanln(&city)
	fmt.Println("Laster inn data fra flere kilder...")
	data := rest.Loaddata(city, "127.0.0.1")
	var weather map[string]float64
	weather = data.Weather.Main
	fmt.Printf(weather[])
	fmt.Println("Lokasjons informasjon:")
	fmt.Println(data.Wiki)
}
