package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"sync"
	"time"
)

func main() {
	// regular expression for an ipv4 address
	reIpv4 := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)

	// use arp to find local addresses
	fmt.Println("Discovering local addresses ...")
	out, err := exec.Command("arp", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}

	// see which addresses are open on port 22
	fmt.Println("Testing ssh connectivity ...")
	var wg sync.WaitGroup
	for _, addr := range strings.Split(string(out), "\n") {
		ip := string(reIpv4.Find([]byte(addr)))
		wg.Add(1)
		go func(ip string) {
			defer func() {
				wg.Done()
			}()
			if _, err := net.DialTimeout("tcp4", ip+":22", 1*time.Second); err == nil {
				log.Printf("connected to %s at port 22", ip)
			}
		}(ip)
	}
	wg.Wait()
}
