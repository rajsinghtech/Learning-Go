package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <ip-address>\n", os.Args[0])
		os.Exit(1)
	}

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Fprintf(os.Stderr, "Invalid IP address: %q\n", os.Args[1])
		os.Exit(1)
	}

	ones, bits := ip.DefaultMask().Size()
	network := ip.Mask(net.CIDRMask(ones, bits))

	for i := 1; i <= 254; i++ {
		addr := net.IPv4(network[0], network[1], network[2], byte(i))
		go ping(addr.String())
	}
}

func ping(ipAddr string) {
	conn, err := net.Dial("ip4:icmp", ipAddr)
	if err != nil {
		fmt.Printf("%s is down\n", ipAddr)
		return
	}
	defer conn.Close()
	fmt.Printf("%s is up\n", ipAddr)
}