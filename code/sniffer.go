package main

import (
	"time"
	"fmt"
	"log"
	"github.com/google/gopacket"
	//"github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

var (	
	snaplen int32 = 1024
	promisc bool  = true
	timeout time.Duration = pcap.BlockForever

	/* prettify */
	cyan   = "\033[36m"
	green  = "\033[32m"
	white  = "\033[37m"
	yellow = "\033[33m"
	red    = "\033[31m"
	fn     = "\033[0m"

	color  = []string{fn,cyan,green,white,yellow,red}
)

func Sniffer(interface_name string) {
	handle, err := pcap.OpenLive(interface_name, snaplen, promisc, timeout)
	
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	
	fmt.Printf("📡 [ MONITORING NETWORK FOR INTERFACE %s ] 📡\n", interface_name)
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
