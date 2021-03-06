package main

//////////////////////////////////
// 	FILEINFO FOR OPPGAVE2	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Skriv filbane på filen du vil scanne:")
	var fileinfo = bufio.NewScanner(os.Stdin)
	var fileinput string = "Ingen fil valgt"
	for fileinfo.Scan() {
		fileinput = string(fileinfo.Text())
		break
	}
	fileInfo(fileinput)
}
func fileInfo(filename string) {
	fmt.Println("Scanner fil......")
	a, err := os.Stat(filename)
	if err != nil {
	}
	//Størrelse i forskjellig former:
	size := float64(a.Size())
	kibi := size / 1000
	mibi := size / 1000000
	gibi := size / 1000000000
	fmt.Println("Bytes: ", size)
	fmt.Println("Kibibytes: ", kibi)
	fmt.Println("Mibibytes: ", mibi)
	fmt.Println("Mibibytes: ", gibi)
	// Switch types mappe eller fil
	switch mode := a.Mode(); {
	case mode.IsRegular():
		fmt.Println("Dette er en fil")
	case mode.IsDir():
		fmt.Println("Dette er en mappe")
	case mode&os.ModeAppend != 0:
		fmt.Println("Dette er en snarvei (symbolic link)")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("Dette er en named pipe")
	}
	fmt.Println("Filrettigheter: ", a.Mode())
	mode := a.Mode()
	if mode&os.ModeAppend == 0 {
		fmt.Println("Filen er ikke 'append only'")
	} else {
		fmt.Println("Filen er 'append only'")
	}
	if mode&os.ModeNamedPipe == 0 {
		fmt.Println("Filen er ikke 'name pipe'")
	} else {
		fmt.Println("Filen er 'name pipe'")
	}
	if mode&os.ModeDevice == 0 {
		fmt.Println("Filen er ikke  'device oppgaver'")
	} else {
		fmt.Println("Filen er 'device oppgaver'")
	}
	if mode&os.ModeCharDevice == 0 {
		fmt.Println("Filen er ikke  ' Unix character device'")
	} else {
		fmt.Println("Filen er ' Unix character device'")
	}
	if mode&os.ModeSymlink == 0 {
		fmt.Println("Filen er ikke  'symbolic link'")
	} else {
		fmt.Println("Filen er 'symbolic link'")
	}
}