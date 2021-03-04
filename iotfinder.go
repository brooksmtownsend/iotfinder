package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func main() {
	fmt.Println("Discovering local addresses ...")

	out, err := exec.Command("arp", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}

	// regular expression for an ipv4 address
	reIpv4 := regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+`)

	fmt.Println("Testing ssh connectivity ...")
	for _, addr := range strings.Split(string(out), "\n") {
		//TODO(brooksmtownsend): append port to ip address, dial
		ip := string(reIpv4.Find([]byte(addr)))
		s, _ := time.ParseDuration("1s")
		_, err = net.DialTimeout("tcp4", ip+"22", s)
		if err == nil {
			log.Printf("connected to %s at port 22", ip)
		}
	}

}

// tcpAddr := net.TCPAddr{IP: net.ParseIP(string(reIpv4.Find([]byte(addr)))), Port: 22}
