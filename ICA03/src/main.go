package main

//////////////////////////////////
// 	ICA03, IS-105	 //Zwirc//
//////////////////////////////////

import (
	"./ascii"
	"./iso"
	"./convert"
	"./fileutils"
	"./treasure"
	"bufio"
	"os"
	"fmt"
	"strings"
	"golang.org/x/text/encoding/charmap"

	"encoding/hex"
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
		case "2c":
			fmt.Println("Oppgave 2c:");
			fmt.Println("Kjør egen test av filene i testmappen")
			ferdig();
			break
		case "3a":
			fmt.Println("Oppgave 3a:");
			oppgave3a();
			ferdig();
			break
		case "3b":
			fmt.Println("Oppgave 3b:");
			oppgave3b();
			ferdig();
			break
		case "3b2":
			fmt.Println("Oppgave 3b2:");
			oppgave3b2();
			ferdig();
			break
		case "3c":
			fmt.Println("Oppgave 3c:");
			oppgave3c();
			ferdig();
			break
		case "4a":
			fmt.Println("Oppgave 4a:");
			oppgave4a();
			ferdig();
			break
		case "4b":
			fmt.Println("Oppgave 4b:");
			oppgave4b();
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
		"2b = Oppgave 2b\n" +
		"2c = Oppgave 2c\n" +
		"3a = Oppgave 3a\n" +
		"3b = Oppgave 3b\n" +
		"3b2 = Oppgave 3b2\n" +
		"3c = Oppgave 3c\n" +
		"4a = Oppgave 4a\n" +
		"4b = Oppgave 4b")
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
	fmt.Println("\nRetur: \n" + iso.GreetingExtendedASCII())
}
func oppgave3a(){
	convert.Utf8()
}
func oppgave3b(){
	byteslice1 := fileutils.FileToByteslice("C:/OneDrive/Skole/Github/ICA/ICA03/src/files/lang01.wl")
	fmt.Println(byteslice1)
	byteslice2 := fileutils.FileToByteslice("C:/OneDrive/Skole/Github/ICA/ICA03/src/files/lang02.wl")
	fmt.Println(byteslice2)
	byteslice3 := fileutils.FileToByteslice("C:/OneDrive/Skole/Github/ICA/ICA03/src/files/lang03.wl")
	fmt.Println(byteslice3)
	fmt.Println("Skriv 3b2 for symboler")

}
func oppgave3b2(){
	byteslice1 := fileutils.FileToByteslice("C:/OneDrive/Skole/Github/ICA/ICA03/src/files/lang01.wl")
	fmt.Printf("Slice 1: %s \n", []byte(byteslice1))
	byteslice2 := fileutils.FileToByteslice("C:/OneDrive/Skole/Github/ICA/ICA03/src/files/lang02.wl")
	fmt.Printf("Slice 2: %s \n", []byte(byteslice2))
	byteslice3 := fileutils.FileToByteslice("C:/OneDrive/Skole/Github/ICA/ICA03/src/files/lang03.wl")
	fmt.Printf("Slice 3: %s \n", []byte(byteslice3))

}
func oppgave3c() {
	treasureString := "C:/OneDrive/Skole/Github/ICA/ICA03/src/treasure/treasure.txt"
	byteslice := treasure.PrintTreasureUTF8(treasureString)
	fmt.Printf("Slice: %x \n", []byte(byteslice))

	// Use charmap library to convert to KOI8R
	encoded, err := charmap.KOI8R.NewDecoder().Bytes(byteslice)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println("Konvertert HEX")
	fmt.Printf("%s", encoded)

	const result = "\x48\x65\x6e\x72\x69\x6b\x20\x41\x72\x6e\x6f\x6c\x64\x20\x57\x65\x72\x67\x65\x6c\x61\x6e\x64\x20\x28\x66\xf8\x64\x74\x20\x31\x37\x2e\x20\x6a\x75\x6e\x69\x20\x31\x38\x30\x38\x2c\x20\x64\xf8\x64\x20\x31\x32\x2e\x20\x6a\x75\x6c\x69\x20\x31\x38\x34\x35\x29\x0a\x56\x69\x20\x65\x72\x65\x20\x65\x6e\x20\x6e\x61\x73\x6a\x6f\x6e\x20\x76\x69\x20\x6d\x65\x64\x2c\x0a\x20\x76\x69\x20\x73\x6d\xe5\x20\x65\x6e\x20\x61\x6c\x65\x6e\x20\x6c\x61\x6e\x67\x65\x2c\x0a\x20\x65\x74\x20\x66\x65\x64\x72\x65\x6c\x61\x6e\x64\x20\x76\x69\x20\x66\x72\x79\x64\x65\x73\x20\x76\x65\x64\x2c\x0a\x20\x6f\x67\x20\x76\x69\x2c\x20\x76\x69\x20\x65\x72\x65\x20\x6d\x61\x6e\x67\x65\x2e\x0a\x20\x56\xe5\x72\x74\x20\x68\x6a\x65\x72\x74\x65\x20\x76\x65\x74\x2c\x20\x76\xe5\x72\x74\x20\xf8\x79\x65\x20\x73\x65\x72\x0a\x20\x68\x76\x6f\x72\x20\x67\x6f\x64\x74\x20\x6f\x67\x20\x76\x61\x6b\x6b\x65\x72\x74\x20\x4e\x6f\x72\x67\x65\x20\x65\x72\x2c\x0a\x20\x76\xe5\x72\x20\x74\x75\x6e\x67\x65\x20\x6b\x61\x6e\x20\x65\x6e\x20\x73\x61\x6e\x67\x20\x62\x6c\x61\x6e\x74\x20\x66\x6c\x65\x72\x0a\x20\x61\x76\x20\x4e\x6f\x72\x67\x65\x73\x20\xe6\x72\x65\x73\x2d\x73\x61\x6e\x67\x65\x2e\x0a\x0a\x20\x4d\x65\x72\x20\x67\x72\xf8\x6e\x74\x20\x65\x72\x20\x67\x72\x65\x73\x73\x65\x74\x20\x69\x6e\x67\x65\x6e\x73\x74\x65\x64\x73\x2c\x0a\x20\x6d\x65\x72\x20\x66\x75\x6c\x6c\x74\x20\x61\x76\x20\x62\x6c\x6f\x6d\x73\x74\x65\x72\x20\x76\x65\x76\x65\x74\x0a\x20\x65\x6e\x6e\x20\x69\x20\x64\x65\x74\x20\x6c\x61\x6e\x64\x20\x68\x76\x6f\x72\x20\x6a\x65\x67\x20\x74\x69\x6c\x66\x72\x65\x64\x73\x0a\x20\x6d\x65\x64\x20\x66\x61\x72\x20\x6f\x67\x20\x6d\x6f\x72\x20\x68\x61\x72\x20\x6c\x65\x76\x65\x74\x2e\x0a\x20\x4a\x65\x67\x20\x76\x69\x6c\x20\x64\x65\x74\x20\x65\x6c\x73\x6b\x65\x20\x74\x69\x6c\x20\x6d\x69\x6e\x20\x64\xf8\x64\x2c\x0a\x20\x65\x69\x20\x62\x79\x74\x74\x65\x20\x64\x65\x74\x20\x68\x76\x6f\x72\x20\x6a\x65\x67\x20\x65\x72\x20\x66\xf8\x64\x64\x2c\x0a\x20\x6f\x6d\x20\x6d\x61\x6e\x20\x65\x74\x20\x70\x61\x72\x61\x64\x69\x73\x20\x6d\x65\x67\x20\x62\xf8\x64\x0a\x20\x61\x76\x20\x70\x61\x6c\x6d\x65\x72\x20\x6f\x76\x65\x72\x73\x76\x65\x76\x65\x74\x2e\x0a"
	fmt.Printf("%s", result)

}
func oppgave4a() {
	fmt.Println("Error ved å dekode til string:")
	fmt.Println(hex.DecodeString("“nord og sør” på islandsk er “norður og suður”"))

}
func oppgave4b() {
	fmt.Println("Kjørt i nettleser...")
}