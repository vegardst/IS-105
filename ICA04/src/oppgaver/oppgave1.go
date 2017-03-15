package oppgaver

//////////////////////////////////
// 	ICA04, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////
import (
	"os"
	"log"
	"regexp"
	"io"
	"io/ioutil"
	"fmt"
	"bufio"
)

// For oppgave 1a
// Lik metode som i forige ICA, som leser fil og returnerer den som byteslice
func Oppgave1a(source string) {
	// Bruker samme metode som vi fikk i ICA03 for å gjøre om text filene til byteslice's
	text1 := fileToByteslice(source + "files/text1.txt")
	text2 := fileToByteslice(source + "files/text2.txt")
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
}
func Oppgave1b(source string) {
	text1 := source + "files/text1.txt"
	text2 := source + "files/text2.txt"
	// Sjekker for linjeskift i oppgaver klassen.
	fmt.Println("Text 1: ", checkLineshift(text1))
	fmt.Println("Text 2: ", checkLineshift(text2))

	// Starter en ny scanner for å få input av filbane.
	fmt.Println("Skriv filbane på filen du vil scanne: (Root soruce definert i var)")
	var filescanner = bufio.NewScanner(os.Stdin)
	var fileinput string = "Ingen fil valgt"
	for filescanner.Scan() {
		fileinput = source + string(filescanner.Text())
		break
	}
	fmt.Println("Din fil:", fileinput)
	fmt.Println("Scanning : ", checkLineshift(fileinput))

}

func fileToByteslice(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()
	byteSlice := make([]byte, size_of_slice)
	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	return byteSlice
}

// For oppgave 1 b
// Søker igjennom string for linjeskift
func checkLineshift(filename string) string {
	textLF := FindLindebreak(filename)
	textCF := FindCarriageReturn(filename)
	var lf bool = false
	var cf bool = false
	// Bruker boolean for å indikere linjeskift LF og CF
	if len(textLF) > 0 {
		lf = true
	}
	if len(textCF) > 0 {
		cf = true
	}
	if (lf && cf) {
		return "Teksten inneholder LineBreak og CarriageReturn"
	} else if (lf) {
		return "Teksten inneholder LineBreak "
	} else if (cf) {
		return "Teksten inneholder CarriageReturn"
	}
	return "Teksten inneholder ikke linjeskift."
}

// Finner Line Break i fil
func FindLindebreak(filename string) []byte {
	var lf = regexp.MustCompile("\n+")
	a, _ := ioutil.ReadFile(filename)
	return lf.Find(a)
}

// Finner Carriage Return i fil
func FindCarriageReturn(filename string) []byte {
	var lf = regexp.MustCompile("\r+")
	b, _ := ioutil.ReadFile(filename)
	return lf.Find(b)
}
