package generators

import (
	"fmt"
	"github.com/timoisik/web-map/models"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GenerateDomainsByCrawling(domainsChannel chan models.Domain) {
	var domains []models.Domain
	models.Db.Where("crawled_at is null or crawled_at = ?", "0001-01-01 00:00:00+00").Find(&domains)

	for _, domain := range domains {
		crawlDomain(domainsChannel, domain)
	}
}

func crawlDomain(domainsChannel chan models.Domain, domain models.Domain) {
	fmt.Printf("Crawling: %v\n", domain.GetUrl())

	url := fmt.Sprintf("http://%v", domain.GetUrl())
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Cannot GET url %v\n", domain.GetUrl())
		return
	}

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Cannot read response Body for url %v\n", domain.GetUrl())
	}

	domain.MarkAsCrawled()

	re := regexp.MustCompile(`(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	foundDomains := re.FindAllString(string(html), -1)

	for _, foundDomain := range foundDomains {
		regexDomain := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)\.([^:\/\n]+)`)
		submatchall := regexDomain.FindStringSubmatch(foundDomain)

		newDomain := models.Domain{Name: submatchall[1], Tld: submatchall[2]}
		if !newDomain.Exists() {
			domainsChannel <- newDomain
		}

		if !newDomain.HasBeenCrawled() {
			go crawlDomain(domainsChannel, newDomain)
		}
	}
}