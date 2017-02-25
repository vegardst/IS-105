package main

//////////////////////////////////
// 	ICA03, IS-105	 //Zwirc//
//////////////////////////////////

import (
	"./ascii"
	"./iso"
	"bufio"
	"os"
	"fmt"
	"strings"
)
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	menu()
	for scanner.Scan() {
		input := strings.ToLower(scanner.Text())
		switch input{
		case "hjelp":
			menu()
		case "exit":
			return
		case "1a":
			fmt.Println("Oppgave 1a:");
			oppgave1a();
			ferdig();
			break
		case "1b":
			fmt.Println("Oppgave 1b:");
			oppgave1b();
			ferdig();
			break
		case "1c":
			fmt.Println("Oppgave 1c:");
			oppgave1c();
			ferdig();
			break
		case "2a":
			fmt.Println("Oppgave 2a:");
			oppgave2a();
			ferdig();
			break
		case "2b":
			fmt.Println("Oppgave 2b:");
			oppgave2b();
			ferdig();
			break
		default:
			fmt.Printf("Du må velge nummer på type!")
		}
	}
}

func menu(){
	fmt.Println("Skriv nummer på Oppgave!\nAlternativer:\n \n" +
		"hjelp = Viser menyen \n" +
		"exit = Avslutter programmet \n\n" +
		"1a = Oppgave 1a \n" +
		"1b = Oppgave 1b\n" +
		"1c = Oppgave 1c\n" +
		"2a = Oppgave 2a\n" +
		"2b = Oppgave 2b")
}
func ferdig(){
	fmt.Println("Skriv neste oppgave eller hjelp for meny og exit for å avslutte!")
}
func oppgave1a(){
	ascii.IterateOverASCIIStringLiteral(ascii.Ascii);
}
func oppgave1b(){
	print := ascii.GreetingASCII()
	fmt.Println("Print av return HEX:")
	fmt.Println(print)
}
func oppgave1c(){
	isAssci := ascii.Ascii
	fmt.Println(isAssci)
	if ascii.IsASCII(isAssci){
		fmt.Println(" = Koden er ASCII!")
	}
	if !ascii.IsASCII(isAssci) {
		fmt.Println(" = Koden er ikke ASCII!")
	}

	isAssci2 := "ØASÅ"
	fmt.Print(isAssci2)
	if ascii.IsASCII(isAssci2){
		fmt.Println(" = Koden er ASCII!")
	}
	if !ascii.IsASCII(isAssci2) {
		fmt.Println(" = Koden er ikke ASCII!")
	}
	fmt.Println("Se egen ascii_test.go fil for selve Test oppgaven")
}

func oppgave2a(){
	iso.CreateExtendedAscii()
}

func oppgave2b(){
	fmt.Println("Retur: \n" + iso.GreetingExtendedASCII())
}