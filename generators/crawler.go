package generators

import (
	"fmt"
	"github.com/timoisik/web-map/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func GenerateDomainsByCrawling(domainsChannel chan models.Domain) {
	var domains []models.Domain
	models.Db.Where("crawled_at is null").Find(&domains)

	for _, domain := range domains {
		crawlDomain(domainsChannel, domain)
	}
}

func crawlDomain(domainsChannel chan models.Domain, domain models.Domain) {
	fmt.Printf("Crawling: %v\n", domain.GetUrl())

	url := fmt.Sprintf("http://%v", domain.GetUrl())
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	domain.MarkAsCrawled()

	re := regexp.MustCompile(`(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	foundDomains := re.FindAllString(string(html), -1)

	for _, foundDomain := range foundDomains {
		regexDomain := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)\.([^:\/\n]+)`)
		submatchall := regexDomain.FindStringSubmatch(foundDomain)

		var newDomain models.Domain
		models.Db.FirstOrInit(&newDomain, models.Domain{Name: submatchall[1], Tld: submatchall[2]})

		if !newDomain.HasBeenCrawled() {
			domainsChannel <- newDomain
			crawlDomain(domainsChannel, newDomain)
		}
	}
}