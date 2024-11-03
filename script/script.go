package main

import (
	"fmt"
	"net"
	"os/exec"
	"sync"
)

func ping(ip string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Execute ping command
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip) // Use "-n" for Windows
	if err := cmd.Run(); err == nil {
		fmt.Println(ip, "is active")
	}
}

func main() {
	var subnet string
	fmt.Print("Enter the subnet (e.g., 192.168.1.0/24): ")
	fmt.Scan(&subnet)

	// Parse the subnet
	_, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		fmt.Println("Invalid subnet:", err)
		return
	}

	var wg sync.WaitGroup

	// Loop through all IP addresses in the subnet
	for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); increment(ip) {
		wg.Add(1)
		go ping(ip.String(), &wg)
	}

	wg.Wait()
}

// increment increments the given IP address by 1
func increment(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] != 0 {
			break
		}
	}
}
