package generators

//var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
//	"t", "u", "v", "w", "x", "y", "z"}

var alphabet = []string{"a", "b", "c"} // Using only three letters for testing

var tlds = [...]string{"de"}

var maxDomainLength = 3 // 5 for production and 3 for testing

func GenerateDomains(domainsChannel chan string) {
	combinations("", domainsChannel)
	close(domainsChannel)
}

func combinations(prefix string, ch chan string) {
	for i := 0; i < len(alphabet); i++  {
		domain := prefix + alphabet[i]

		if len(domain) > maxDomainLength {
			return
		}

		for j := 0; j < len(tlds); j++ {
			ch <- domain + "." + tlds[j]
		}

		combinations(domain, ch)
	}
}