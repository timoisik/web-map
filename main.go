package main

import (
	"github.com/timoisik/web-map/fetchers"
	"github.com/timoisik/web-map/generators"
)

func main() {
	domainsChannel := make(chan string)
	go generators.GenerateDomains(domainsChannel)

	for domain := range domainsChannel {
		fetchers.FetchDomain(domain)
	}
}