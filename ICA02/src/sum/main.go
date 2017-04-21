package main
//////////////////////////////////
// 	ICA02, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)


func main() {
	menu()
	for scanner.Scan() {
		input := strings.ToUpper(scanner.Text())
		switch input{
		case "0":
			menu()
		case "1": // int8
			numbers("int8");
			break
		case "2": // int8
			numbers("int32");
			break
		case "3": // int8
			numbers("int64");
			break
		case "4": // int8
			numbers("uint32");
			break
		case "5": // int8
			numbers("float64");
			break
		default:
			fmt.Printf("Du må velge nummer på type!")
		}
		break
	}
}

func menu(){
	fmt.Println("Velg nummer for utregnings metode:\n" +
		"0. Vis meny	\n" +
		"1. int8 	\n" +
		"2. int32	\n" +
		"3. int64	\n" +
		"4. uint32	\n" +
		"5. float64	")
}

func numbers(format string) {
	fmt.Println("Skriv inn første " + format  + " tall")
	var number1 string
	var number2 string
	for scanner.Scan() {
		number1 = strings.ToUpper(scanner.Text())
		break
	}
	fmt.Println("Skriv inn neste " + format + " tall ")
	for scanner.Scan() {
		number2 = strings.ToUpper(scanner.Text())
		break
	}
	checkFormat(format, number1, number2)
}

func checkFormat(format string, number1 string, number2 string){
	var value float64 = checkFlow(number1, number2)
	switch format{
		case "int8":
			// Try to convert numbers to int8, base 10.
			n1, err := strconv.ParseInt(number1, 10, 8)
			n2, err := strconv.ParseInt(number2, 10, 8)
			// Set value as int8
			a := int8(n1)
			b := int8(n2)
			// Print error on input
			if err != nil {
				fmt.Println("Dine tall er ikke en int8 (Int8 er 8-bit i range -128 to 127")
				break
			}
			// Check value for OVERFLOW
			if value > 127 {
				fmt.Println("Resultat OVERFLOW! Bruk annen type!")
				break
			}
			// Check value for UNDERFLOW
			if value < -128 {
				fmt.Println("Resultat UNDERFLOW! Bruk annen type!")
				break
			}
			// Print result if everything is ok!
			fmt.Println("Ditt resultat: " , SumInt8(a,b))
		// End application
		break

	case "int32":
		n1, err := strconv.ParseInt(number1, 10, 32)
		n2, err := strconv.ParseInt(number2, 10, 32)
		a := int32(n1)
		b := int32(n2)
		if err != nil {
			fmt.Println("Dine tall er ikke en int32 (Int32 er 32-bit i range -2147483648 to 2147483647")
			break
		}
		if value > 2147483647 {
			fmt.Println("Resultat OVERFLOW! Bruk annen type!")
			break
		}
		if value < -2147483648 {
			fmt.Println("Resultat UNDERFLOW! Bruk annen type!")
			break
		}
		fmt.Println("Ditt resultat: " , SumInt32(a,b))
		break

	case "int64":
		n1, err := strconv.ParseInt(number1, 10, 64)
		n2, err := strconv.ParseInt(number2, 10, 64)
		a := int64(n1)
		b := int64(n2)
		if err != nil {
			fmt.Println("Dine tall er ikke en int64 (int64 er 64-bit i range -9223372036854775808 to 9223372036854775807")
			break
		}
		if value > 9223372036854775807 {
			fmt.Println("Resultat OVERFLOW! Bruk annen type!")
			break
		}
		if value < -9223372036854775808 {
			fmt.Println("Resultat UNDERFLOW! Bruk annen type!")
			break
		}
		fmt.Println("Ditt resultat: " , SumInt64(a,b))
		break

	case "uint32":
		n1, err := strconv.ParseUint(number1, 10, 32)
		n2, err := strconv.ParseUint(number2, 10, 32)
		a := uint32(n1)
		b := uint32(n2)
		if err != nil {
			fmt.Println("Dine tall er ikke en uint32 (uint32 er 32-bit i range 0 to 4294967295")
			break
		}
		if value > 4294967295 {
			fmt.Println("Resultat OVERFLOW! Bruk annen type!")
			break
		}
		if value < 0 {
			fmt.Println("Resultat UNDERFLOW! Bruk annen type!")
			break
		}
		fmt.Println("Ditt resultat: " , SumUint32(a,b))
		break

	case "float64":
		n1, err := strconv.ParseFloat(number1, 64)
		n2, err := strconv.ParseFloat(number2, 64)
		a := float64(n1)
		b := float64(n2)
		if err != nil {
			fmt.Println("Dine tall er ikke en uint32 (uint32 er 32-bit i range -9223372036854775808 to 9223372036854775807")
			break
		}
		fmt.Println("Float64 er av veldig høy type, så ditt resultat blir ikke over/under-flow sjekket.")
		fmt.Println("Ditt resultat: " , SumFloat64(a,b))
		break
	}
}

func checkFlow(number1 string, number2 string) float64 {
	a, err := strconv.ParseFloat(number1, 64)
	b, err := strconv.ParseFloat(number2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return a+b
}
