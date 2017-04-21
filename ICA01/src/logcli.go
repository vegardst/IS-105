package main

import "os"
import "fmt"
import "math"
import "strconv"

func main() {
    arg := os.Args[1]
    // Bruker strconv (String Converter) fra bibloteket for å få float64
    input, err := strconv.ParseFloat(arg, 64)
      if err != nil {
        fmt.Println("Error! Brukte du riktig tall format?")
    }
    fmt.Println("Base-2 utregning av dit tall: " +arg+ "blir:")
    fmt.Println(math.Log2(input))
    fmt.Println("Base-10 utregning av dit tall: " +arg+ "blir:")
    fmt.Println(math.Log10(input))
}
