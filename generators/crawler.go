package generators

import (
	"fmt"
	"github.com/timoisik/web-map/models"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

func GenerateDomainsByCrawling(domainsChannel chan models.Domain) {
	var domains []models.Domain
	models.Db.Where("crawled_at is null or crawled_at = ?", "0001-01-01 00:00:00+00").Find(&domains)

	var wg sync.WaitGroup

	for _, domain := range domains {
		wg.Add(1)
		crawlDomain(domainsChannel, &wg, domain)
	}

	wg.Wait()
	close(domainsChannel)
}

func crawlDomain(domainsChannel chan models.Domain, wg *sync.WaitGroup, domain models.Domain) {
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

	re := regexp.MustCompile(`(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	foundDomains := re.FindAllString(string(html), -1)

	domain.MarkAsCrawled()

	for _, foundDomain := range foundDomains {
		regexDomain := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)\.([^:\/\n]+)`)
		submatchall := regexDomain.FindStringSubmatch(foundDomain)

		newDomain := models.Domain{Name: submatchall[1], Tld: submatchall[2]}
		if !newDomain.Exists() {
			domainsChannel <- newDomain
		}

		if !newDomain.HasBeenCrawled() {
			wg.Add(1)
			go crawlDomain(domainsChannel, wg, newDomain)
		}
	}

	wg.Done()
}