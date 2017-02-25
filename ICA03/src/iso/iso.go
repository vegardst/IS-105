package iso

import (
	"fmt"
	"encoding/json"
)

//Oppgave 2A
func CreateExtendedAscii(){
	var ascii []byte
	for i := 0x80; i <= 0xFF; i++ {
		ascii = append(ascii, byte(i))
	}
	IterateOverASCIIStringLiteral(ascii)
}

func IterateOverASCIIStringLiteral(ascii []byte) {
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X %q %b \n", ascii[i], ascii[i], ascii[i])
	}

	// Alternativ for å se data fra bytes:
	fmt.Println("Alternativ print-string av byte listen:")
	data := json.RawMessage(ascii)
	bytes, err := data.MarshalJSON()
	if err != nil {
		fmt.Println("error")
	} else {
		//print console
		fmt.Println(string(bytes))
	}
}

//Oppgave 2B
func GreetingExtendedASCII() string{
	greetings := []byte{83, 97, 108, 117, 116, 44, 32, 231, 97, 32, 118, 97, 32, 176, 45, 41, 32, 128, 53, 48}
	for i := 0; i < len(greetings); i++ {
		fmt.Printf("%c", greetings[i])
	}
	retur := string(greetings)
	return retur
}


