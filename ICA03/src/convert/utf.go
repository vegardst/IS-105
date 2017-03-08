package convert

import "fmt"

const a = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
const b = "\xC2\xBD\x3F\x3D\x3F\x20\xE2\x8C\x98"

func Utf8() {
	fmt.Printf("%%s: %s\n", a)
	fmt.Printf("%%q: %q\n", a)
	fmt.Printf("%%+q: %+q\n", a)
	fmt.Printf("%%X: %X\n", a)
	fmt.Printf("%% X: % X\n", a)
	fmt.Printf("%%c: %c\n", []byte(a))
	fmt.Printf("1/2?=? ⌘ = %s \n", []byte(b))
}


