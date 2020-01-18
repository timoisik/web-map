package fetchers

import (
	"fmt"
	"net"
)

func FetchDomain(domain string) {
	ip, err := net.LookupIP(domain)

	if err != nil {
		fmt.Printf("No ip for domain %v\n", domain)
		return
	}

	fmt.Printf("Domain %v exists at %v\n", domain, ip)
}