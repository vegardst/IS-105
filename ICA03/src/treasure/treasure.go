package treasure

import (
	"os"
	"log"
	"io"
	"io/ioutil"
	"strings"
	"strconv"
	"golang.org/x/text/encoding/charmap"
)

// Kode for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasure(filename string)[]byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	return byteSlice // returverdien er her kun en stedsholder
}
func PrintTreasureUTF8(filename string) string {
	var byteslice []byte
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range strings.Split(string(input), "\\x") {
		if (len(i) > 0) {
			i = strings.TrimSpace(i)
			i, _ := strconv.ParseUint(i, 16, 16)
		byteslice = append(byteslice, byte(i))
		}
	}
	returnstring, _ := charmap.ISO8859_1.NewDecoder().Bytes(byteslice)
	return string(returnstring)
}