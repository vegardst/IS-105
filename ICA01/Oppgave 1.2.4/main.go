// Kjør filen ved "go run main.go log.go
// eller bygg filen "go build ..."
package main

import (
	"fmt"
  "./log"
)
func main() {
	fmt.Println("Hello, playground")
	// Kalle public Oppgave i Log.go i samme mappe:
  Oppgave()
	// Log under ./log mappen:
	log.Oppgave()

}
