package main

import (
	"fmt"
	"log"
	"os/exec"

)

func main() {
	os()
	cpu()
	memcache()
	ram()

}

func cpu(){
	out, err := exec.Command("Wmic", "cpu", "get", "Name" , ",Description" , ",SocketDesignation").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CPU INFO: Architecture: Haswell x86 \n %s", out)

}

func memcache(){
	mem, err := exec.Command("Wmic", "memcache" , "get", "Purpose", ",NumberOfBlocks").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Memory Cache info:\n %s", mem)
}

func ram(){
	mem, err := exec.Command("wmic", "MemoryChip", "get" ,"ConfiguredClockSpeed", ",Capacity" , ",Tag").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Memory Cache info:\n %s", mem)

}

func os(){
	out, err := exec.Command("Wmic", "os" , "get", "Caption", ).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OS Info \n %s", out)
}