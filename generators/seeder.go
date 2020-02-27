package generators

import (
	"bufio"
	"github.com/timoisik/web-map/models"
	"log"
	"os"
	"strings"
)

func GenerateBySeeder(domainsChannel chan models.Domain, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domainsChannel <- models.Domain{Name: strings.ToLower(scanner.Text()), Tld: "de"}
	}
}