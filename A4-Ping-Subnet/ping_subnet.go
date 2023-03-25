package main

import (
    "fmt"
    "net"
    "time"
	"os/exec"
)

func main() {
    // Replace with your subnet in CIDR notation
    subnet := "192.168.1.0/24"

    // Parse the subnet
    ip, ipnet, err := net.ParseCIDR(subnet)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Loop through all the IP addresses in the subnet
    for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
        go func(ip net.IP) {
            // Ping the IP address
            if err := ping(ip.String()); err != nil {
                fmt.Printf("Ping failed for %s: %s\n", ip, err)
            } else {
                fmt.Printf("%s is up\n", ip)
            }
        }(ip)
    }

    // Wait for pings to finish
    time.Sleep(10 * time.Second)
}

func ping(ip string) error {
    cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
    return cmd.Run()
}

func inc(ip net.IP) {
    for j := len(ip) - 1; j >= 0; j-- {
        ip[j]++
        if ip[j] > 0 {
            break
        }
    }
}