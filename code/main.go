package main

import (
	"os"
	"fmt"
)

func usage() {
	fmt.Println("usage: ./dog <NETWORK INTERFACE NAME>")
}

func main() {
	if (len(os.Args) != 2) {
		usage()
	}
	Sniffer(string(os.Args[1]))
}
