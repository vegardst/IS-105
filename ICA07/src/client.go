package main

import "net"
import "fmt"
import "bufio"
import "os"

const (
	HOST =  "127.0.0.1"
	PORT = "8001"
)

func main() {
	fmt.Println("Vil du sende TCP eller UDP?")
	var stype string = "TCP"
	fmt.Scanln(&stype)
	if (stype == "UDP" || stype == "udp" ) {
		send("udp")
	} else if (stype == "TCP" || stype == "tcp" ) {
		send("tcp")
	} else {
		fmt.Println("Kunne ikke gjenkjenne type, TCP satt som standard.")
		send("tcp")
	}
}
func send(stype string){
	conn, _ := net.Dial(stype, HOST+":"+PORT)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Skriv melding til serveren: ")
		text, _ := reader.ReadString('\n')
		cryptTxt := crypt(text)
		fmt.Fprintf(conn, string(cryptTxt))
		fmt.Println("Tekst kryptert og sendt! Melding:" + text)
	}
}