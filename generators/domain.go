package generators

import (
	"github.com/timoisik/web-map/models"
	"time"
)

var characters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z"}

//var characters = []string{"a", "b", "c"} // Using only three letters for testing

var tlds = [...]string{"de"}

var maxDomainLength = 5 // 5 for production and 3 for testing

func GenerateDomains(domainsChannel chan models.Domain) {
	combinations("", domainsChannel)
	close(domainsChannel)
}

func combinations(prefix string, ch chan models.Domain) {
	for i := 0; i < len(characters); i++  {
		domainName := prefix + characters[i]

		if len(domainName) > maxDomainLength {
			return
		}

		for j := 0; j < len(tlds); j++ {
			ch <- models.Domain{Name: domainName, Tld: tlds[j], CreatedAt: time.Now()}
		}

		combinations(domainName, ch)
	}
}