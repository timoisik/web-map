package generators

var alphabet = []string{"a", "b", "c"}

var tlds = [...]string{"de"}

func GenerateDomains(domainsChannel chan string) {
	combinations("", 1, domainsChannel)
	close(domainsChannel)
}

func combinations(prefix string, length int, ch chan string) {

	for i := 0; i < len(alphabet); i++  {
		domain := prefix + alphabet[i]

		if length > len(alphabet){
			return
		}

		for j := 0; j < len(tlds); j++ {
			ch <- domain + "." + tlds[j]
		}

		combinations(domain, length + 1, ch)
	}
}