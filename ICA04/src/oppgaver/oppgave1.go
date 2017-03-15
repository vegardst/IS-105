package oppgaver

import (
	"os"
	"log"
	"regexp"
	"io"
	"io/ioutil"
)

// For oppgave 1a
// Lik metode som i forige ICA, som leser fil og returnerer den som byteslice
func FileToByteslice(filename string) []byte {
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
func CheckLineshift(filename string) string {
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
