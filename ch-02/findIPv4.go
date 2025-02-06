package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"regexp"
)

//the regular expression defined in partIP matches each
//one of the four parts of an IP address
//A valid Ip address can begin with 25 and end with
//0, 1, 3, 4, or 5
//because that's the biggest 8-bit binary number(25[0-5])
//or it can begin with 2 followed by 0, 1, 2, 3, 4, and end with
//0, 1, 2, 3, 4, 5, 6, 7, 8, or 9(2[0-4][0-9]).
//Alternatively, it can begin with 1 followed by two or more
//digits from 0 to 9 (1[0-9][0-9]).
//The last alternative would be a natural number that has one
//or two digits
//The first digit, which is optional, can be from 1 to 9
//and the second, which is mandatory can be from 0 - 9

//The grammar variable tells us that we are looking for has
//four distinct parts and one of them must match partIP.
//That grammar variable is what matches the complete IPV4
//address that we seek.

func findIP(input string) string {
	partIP := "(25[0-5]| 2[0-4][0-9]|1[0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIp + "\\." + partIp + "\\."
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	//assert command line arguments
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}
	//iterate all over all command-line argument
	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening files %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()

		//fourth portion of code
		r := buffio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				btrak
			}
			//The last part
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
		}
	}
}
