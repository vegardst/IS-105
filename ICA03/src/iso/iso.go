package iso

import (
	"fmt"
//	"encoding/json"
)

//Oppgave 2A
func CreateExtendedAscii(){
	var extascii []byte
	for i := 0x80; i <= 0xFF; i++ {
		extascii = append(extascii, byte(i))
	}
	IterateOverASCIIStringLiteral(extascii)
}

func IterateOverASCIIStringLiteral(extascii []byte) {
	for i := 0; i < len(extascii); i++ {
		fmt.Printf("%X %q %b \n", extascii[i], extascii[i], extascii[i])
	}
	fmt.Println("")
	s := string(extascii[:])

	// Print symbol fra byte
	fmt.Println("Byte utskrift:")
	fmt.Println(extascii)
	fmt.Println("")

	// Print string 2
	fmt.Println("HEX av string::")
	fmt.Printf("%q" ,string(s))
	fmt.Println("")
	fmt.Println("")

	// Print string
	fmt.Println("Utskrift i string:")
	fmt.Println(s)
	fmt.Println("")

	// Print string
	fmt.Println("Utskrift i %s av string:")
	fmt.Printf("%s", s)
	fmt.Println("")
	fmt.Println("")



	// Print strnig 3
	//fmt.Println("String eksempel 3:")
	//data := json.RawMessage(extascii)
	//bytes, err := data.MarshalJSON()
	//if err != nil {
	//	fmt.Println("error")
	//} else {
	//	//print console
	//	fmt.Println(string(bytes))
	//}
	fmt.Println("Hvis du ikke ser symbolene, må du endre encoding på din terminal")
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


