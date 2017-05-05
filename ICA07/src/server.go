package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

const (
	HOST = "127.0.0.1"
	PORT = "8001"
)

func main() {
	fmt.Println("Skriv type du vil bruke, UDP eller TCP.")
	var stype string = "TCP"
	fmt.Scanln(&stype)
	if (stype == "UDP" || stype == "udp" ) {
		startUdp()
	} else if (stype == "TCP" || stype == "tcp" ) {
		startTcp()
	} else {
		fmt.Println("Kunne ikke gjenkjenne type, TCP satt som standard.")
		startTcp()
	}
}


//Ongoing 2
func startTcp(){
	fmt.Println("Starter TCP server på " +HOST+":"+PORT)
	// Sett opp adresse og port
	TCP,err := net.ResolveTCPAddr("tcp",HOST+":"+PORT)
	ErrorCheck(err)

	// Lytt på port:
	listen, err := net.ListenTCP("tcp", TCP)
	ErrorCheck(err)
	defer listen.Close()


	ErrorCheck(err)
	for {
		conn , err := listen.Accept()
		ErrorCheck(err)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Melding mottatt:", string(message))


	}
}


func startUdp(){
	fmt.Println("Starter UDP server på " +HOST+":"+PORT)
	// Sett opp adresse og port
	UDP,err := net.ResolveUDPAddr("udp",HOST+":"+PORT)
	ErrorCheck(err)

	// Lytt på port:
	listen, err := net.ListenUDP("udp", UDP)
	ErrorCheck(err)
	defer listen.Close()

	buffer := make([]byte, 1024)

	for {
		n,addr,err := listen.ReadFromUDP(buffer)
		message := buffer[0:n]
		decrypted := decrypt(message)
		fmt.Println("Melding mottatt ",decrypted , " fra ",addr)

		ErrorCheck(err)
	}
}

func ErrorCheck(err error) {
	if err  != nil {
		fmt.Println("Feil: " , err)
		// Stop server
		os.Exit(0)
	}
}