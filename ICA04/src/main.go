package main

import (
	"./oppgaver"
	"strings"
	"fmt"
	"bufio"
	"os"
	"path/filepath"
)

//////////////////////////////////
// 	ICA04, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

// Endre source!
const source = "ICA04/src/"

// Import
// "github.com/mozu0/huffman"1
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	menu()
	for scanner.Scan() {
		input := strings.ToLower(scanner.Text())
		switch input {
		case "hjelp":
			menu()
		case "exit":
			return
		case "1a":
			fmt.Println("Oppgave 1a:");
			oppgaver.Oppgave1a(source);
			ferdig();
			break
		case "1b":
			fmt.Println("Oppgave 1b:");
			oppgaver.Oppgave1b(source);
			ferdig();
			break
		case "2a":
			fmt.Println("Oppgave 2a:");
			oppgaver.Oppgave2a(source);
			ferdig();
			break
		case "2b":
			fmt.Println("Oppgave 2b:");
			oppgaver.Oppgave2b();
			ferdig();
			break
		case "2c":
			fmt.Println("Oppgave 2c:");
			oppgaver.Oppgave2c();
			ferdig();
			break
		case "3a":
			fmt.Println("Oppgave 3a:");
			oppgaver.Oppgave3a();
			ferdig();
			break
		case "3b":
			fmt.Println("Oppgave 3b:");
			oppgaver.Oppgave3b(source);
			ferdig();
			break
		case "3c":
			fmt.Println("Oppgave 3c:");
			oppgaver.Oppgave3c();
			ferdig();
			break
		case "4a":
			fmt.Println("Oppgave 4a:");
			oppgaver.Oppgave4a();
			ferdig();
			break
		case "4b":
			fmt.Println("Oppgave 4b:");
			oppgaver.Oppgave4b();
			ferdig();
			break
		case "4c":
			fmt.Println("Oppgave 4c:");
			oppgaver.Oppgave4c();
			ferdig();
			break
		case "4d":
			fmt.Println("Oppgave 4d:");
			oppgaver.Oppgave4d();
			ferdig();
			break
		case "4e":
			fmt.Println("Oppgave 4e:");
			oppgaver.Oppgave4e();
			ferdig();
			break
		default:
			fmt.Printf("Du må skrive oppgave nummer + oppgave. Eksempel: 1a\n")

		}
	}
}
func menu() {
	p, _ := filepath.Abs(" /")
	fmt.Println("Root path:" + p)
	s, _ := filepath.Abs(source)
	fmt.Println("Source path:" + s)
	fmt.Println("Skriv nummer på Oppgave!\nAlternativer:\n \n" +
		"hjelp = Viser menyen \n" +
		"exit = Avslutter programmet \n\n" +
		"1a = Oppgave 1a \n" +
		"1b = Oppgave 1b\n" +
		"2a = Oppgave 2a\n" +
		"2b = Oppgave 2b\n" +
		"2c = Oppgave 2c\n" +
		"3a = Oppgave 3a\n" +
		"3b = Oppgave 3b\n" +
		"3c = Oppgave 3c\n" +
		"4a = Oppgave 4a\n" +
		"4b = Oppgave 4b\n" +
		"4c = Oppgave 4c\n" +
		"4d = Oppgave 4d\n" +
		"4e = Oppgave 4e")
}
func ferdig() {
	fmt.Println("Skriv neste oppgave eller hjelp for meny og exit for å avslutte!")
}
