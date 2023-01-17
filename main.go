package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	var input string
	flag.StringVar(&input, "ip", "", "")
	flag.Parse()

	f, err := os.Create("ip.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ip, ipnet, err := net.ParseCIDR(input)
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		fmt.Printf("%s \n", ip)

		_, err := f.WriteString(ip.String() + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}
}
