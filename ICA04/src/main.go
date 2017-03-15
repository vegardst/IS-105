package main

import (
	"./oppgaver"
	"strings"
	"fmt"
	"bufio"
	"os"
	"syscall"
)
//////////////////////////////////
// 	ICA04, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

// Endre source!
const source = "C:/OneDrive/Skole/ICA/ICA04/src/"
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
			oppgave1a();
			ferdig();
			break
		case "1b":
			fmt.Println("Oppgave 1b:");
			oppgave1b();
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
			oppgave2c();
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
		case "4c":
			fmt.Println("Oppgave 4c:");
			oppgave4c();
			ferdig();
			break
		case "4d":
			fmt.Println("Oppgave 4d:");
			oppgave4d();
			ferdig();
			break
		case "4e":
			fmt.Println("Oppgave 4e:");
			oppgave4e();
			ferdig();
			break
		default:
			fmt.Printf("Du må skrive oppgave nummer + oppgave. Eksempel: 1a\n")

		}
	}
}
func menu() {
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

// Oppgaver //
func oppgave1a() {
	// Bruker samme metode som vi fikk i ICA03 for å gjøre om text filene til byteslice's
	text1 := oppgaver.FileToByteslice(source + "files/text1.txt")
	text2 := oppgaver.FileToByteslice(source + "files/text2.txt")
	// Printer ut innholdet for hver slice
	fmt.Println("Utskrift %%X, base15:")
	fmt.Printf("ByteSlice Text1: %X \n", []byte(text1))
	fmt.Printf("ByteSlice Text2: %X \n", []byte(text2))
	fmt.Println("Utskrivt %%q , characters")
	fmt.Printf("ByteSlice Text1: %q \n", []byte(text1))
	fmt.Printf("ByteSlice Text2: %q \n", []byte(text2))

	// Mer info om filene:
	//fmt.Println("Full print:")
	//for i := 0; i < len(text1); i++ {
	//	fmt.Printf("%X %+q %b \n", text1[i], text1[i], text1[i])
	//}
	//fmt.Println("linebreak")
	//for i := 0; i < len(text2); i++ {
	//	fmt.Printf("%X %+q %b \n", text2[i], text2[i], text2[i])
	//}
}
func oppgave1b() {
	text1 := source + "files/text1.txt"
	text2 := source + "files/text2.txt"
	// Sjekker for linjeskift i oppgaver klassen.
	fmt.Println("Text 1: ", oppgaver.CheckLineshift(text1))
	fmt.Println("Text 2: ", oppgaver.CheckLineshift(text2))

	// Starter en ny scanner for å få input av filbane.
	fmt.Println("Skriv filbane på filen du vil scanne: (Root soruce definert i var)")
	var filescanner = bufio.NewScanner(os.Stdin)
	var fileinput string = "Ingen fil valgt"
	for filescanner.Scan() {
		fileinput = source + string(filescanner.Text())
		break
	}
	fmt.Println("Din fil:", fileinput)
	fmt.Println("Scanning : ", oppgaver.CheckLineshift(fileinput))
}
func oppgave2a() {
	fmt.Println("Skriv filbane på filen du vil scanne: (Root soruce definert i var)")
	var fileinfo = bufio.NewScanner(os.Stdin)
	var fileinput string = "Ingen fil valgt"
	for fileinfo.Scan() {
		fileinput = source + string(fileinfo.Text())
		break
	}
	oppgaver.FileInfo(fileinput)
}
func oppgave2b() {
	stdin := os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	fmt.Println("​/dev/stdin​ :")
	oppgaver.FileInfo(stdin.Name())
	fmt.Println("/dev/ram0 :")
	oppgaver.FileInfo("/dev/ram0")
}
func oppgave2c() {
	fmt.Println("Filene er bygget og ligger i out mappen.")
	fmt.Println("Alternativ: linux, mac, windows")
}
func oppgave3a() {
	fmt.Println("Ved bruk av os og ioutil pakkene kan man behandle filer. " +
		"I disse pakken finner vi metoder som kan gjøre det meste med filer" +
		"os.Create" +
		"os.Remove" +
		"os.Rename" +
		"os.OpenFile" +
		"os.Open" +
		"os.Chmod")
}
func oppgave3b() {
	fmt.Println("Skriv inn filen du ønsker å scanne (root source)")
	var filelenght = bufio.NewScanner(os.Stdin)
	var fileinput string = "Ingen fil valgt"
	for filelenght.Scan() {
		fileinput = source + string(filelenght.Text())
		break
	}

	fmt.Print("Antall linjeskift i filen: ")
	fmt.Println(oppgaver.FindInText(fileinput, "\n"))

	// Henter map av filen
	m := oppgaver.CountFile(fileinput)
	// Sorterer og printer resultat:
	oppgaver.SortAndPrint(m)


}
func oppgave3c() {
	fmt.Println("Kjør benchmark på egne filer")
}
func oppgave4a() {
	/*
	 Helse og Idrettsdag		1420	1829
	 Humanoira og pedagogikk	1182	1525
	 Kunstfag			 293	 420
	 Teknologi og realfag		1337	2166
	 Lærerutdanning			1158	1506
	 Økonomi og samfunnsvitenskap	2398	3098

	 Antall mulige utfall:		7788	10544
	 */
	oppgaver.UIAlist()

}
func oppgave4b() {
	oppgaver.UIAinfo()
	fmt.Println(" Fakulitetet Kunstfag får minst informasjon som sett over.")

}
func oppgave4c() {
	oppgaver.UIAhuffman()


}
func oppgave4d() {
	oppgaver.UIAaverage()
}

func oppgave4e() {

}
