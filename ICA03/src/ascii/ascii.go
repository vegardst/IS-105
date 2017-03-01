package ascii

import (
	"fmt"
)

const Ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"


func IterateOverASCIIStringLiteral(ascii string) {
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X %+q %b \n", ascii[i], ascii[i], ascii[i])
	}
	//      Linje for linje
	//	fmt.Printf("%q ", sl[i])1
	//	fmt.Printf("%b ", sl[i])
	//	.......

	//fmt.Println("Hentet ASCII fra unicode:")
	//for _, ascii16 := range unicode.ASCII_Hex_Digit.R16 {
//	//fmt.Println(v)
	//	for i := ascii16.Lo; i <= ascii16.Hi; i++ {
	//			fmt.Printf("%X %+q %b \n", i, i, i)
	//
	//	}
	//}
}

func GreetingASCII() string {
	// Print DEC
	fmt.Println("DEC:")
	helloDec := []byte{34, 72, 101, 108, 108, 111, 32, 58, 45, 41, 34}
	for i := 0; i < len(helloDec); i++ {
		fmt.Printf("%c", helloDec[i])
	//	fmt.Printf("%s", string(helloDec[i]))
	}
	fmt.Println()
	// Print HEX
	hellohex := "\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"
	return hellohex
}

// From Go Playground
func IsASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}