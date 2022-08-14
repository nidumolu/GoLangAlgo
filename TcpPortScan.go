/*TCP Port Scanner (improved with goroutines)
Scanns concurrently all 65536 ports on a target (really fast)
*/
package main

// usage: go run TcpPortScan.go -h <ipaddress>
// go run TcpPortScan.go -h www.google.com

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

// port scanning using goroutines
func portScan(ip string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	// choose between tcp or udp
	network := "tcp"
	address := ip + ":" + port
	connection, err := net.Dial(network, address)
	// handle errors
	if err != nil {
		return
	}

	fmt.Printf("Port %s is open\n", port)
	connection.Close()
}

func main() {

	// get argument for ip address
	ip := flag.String("h", "", "select IP address to scan")
	// parse argument
	flag.Parse()
	// set slice to store all 65536 port numbers
	var prt []int
	// set slice to sort all 65536 port numbers converted to string
	prtStr := []string{}
	// make integer slice with 65536 slices
	allP := make([]int, 65536)
	// iterate throught all the 65536 slices append them to prt
	for p := range allP {
		prt = append(prt, p)
	}
	// convert the int slice into string slice
	for i := range prt {
		n := prt[i]
		text := strconv.Itoa(n)
		prtStr = append(prtStr, text)
	}

	var wg sync.WaitGroup

	for _, p := range prtStr {
		// if the counter becomes zero, all goroutines blocked on Wait are released
		wg.Add(1)
		// call portScan function and iterate through every port on ip address concurrently
		go portScan(*ip, p, &wg)
	}

	wg.Wait()
}
