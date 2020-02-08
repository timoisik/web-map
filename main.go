package main

import (
	"fmt"
	"github.com/timoisik/web-map/fetchers"
	"github.com/timoisik/web-map/generators"
	"github.com/timoisik/web-map/models"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	models.InitDB()

	domainsChannel := make(chan models.Domain)
	go generators.GenerateDomains(domainsChannel)

	for domain := range domainsChannel {
		domain.Create()

		ip, err := fetchers.FetchDomainIp(domain)
		if err != nil {
			// Remove domain from DB if no IP was found for it
			fmt.Printf("No ip for domain %v\n", domain.GetUrl())
			domain.Delete()
		} else {
			fmt.Printf("Update Domain %v with ips %v\n", domain.GetUrl(), ip)
			// Save IPs in DB
			domain.Update(bson.D{
				{"$set", bson.D{
					{"ips", ip},
				}},
				{"$currentDate", bson.D{
					{"fetchedat", true},
				}},
			})
		}
	}
}