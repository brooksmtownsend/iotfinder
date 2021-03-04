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
	// regular expression for an ipv4 address
	reIpv4 := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)

	fmt.Println("Discovering local addresses ...")
	out, err := exec.Command("arp", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Testing ssh connectivity ...")
	for _, addr := range strings.Split(string(out), "\n") {
		ip := string(reIpv4.Find([]byte(addr)))
		s, _ := time.ParseDuration("1s")
		if _, err := net.DialTimeout("tcp4", ip+":22", s); err == nil {
			log.Printf("connected to %s at port 22", ip)
		}
	}
}
