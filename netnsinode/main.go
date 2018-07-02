package main

import (
	"fmt"
	"regexp"
	"syscall"
)

const (
	netnsfile = "/proc/self/ns/net"
)

func main() {
	for l := 128; ; l *= 2 {
		b := make([]byte, l)
		n, _ := syscall.Readlink(netnsfile, b)
		if n < 0 {
			n = 0
		}
		if n < l {
			link := string(b[0:n])
			fmt.Printf("%s -> %s\n", netnsfile, link)
			re := regexp.MustCompile("net:\\[(.*)\\]")
			submatches := re.FindStringSubmatch(link)
			if len(submatches) >= 1 {
				fmt.Printf("netns inode: %s\n", submatches[1])
			}
			return
		}
	}
}
