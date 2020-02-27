package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timoisik/web-map/generators"
	"github.com/timoisik/web-map/models"
)

// generateDomainCmd represents the createDomain command
var generateDomainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Generates domain",
	Long: `Generates domains for later usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		generateDomains(cmd)
	},
}

func init() {
	generateCmd.AddCommand(generateDomainCmd)

	generateDomainCmd.Flags().BoolP("chars", "c", false, "Generate domains by combination of chars")
	generateDomainCmd.Flags().BoolP("crawl", "r", false, "Generate domains by crawling existing domains")
	generateDomainCmd.Flags().StringP("seed", "s", "", "Generate domain by given seeder list")
}

func generateDomains(cmd *cobra.Command) {
	fmt.Println("generateDomain called")

	chars, _ := cmd.Flags().GetBool("chars")
	crawl, _ := cmd.Flags().GetBool("crawl")
	seed, _ := cmd.Flags().GetString("seed")

	domainsChannel := make(chan models.Domain)

	if chars {
		go generators.GenerateDomainsByChars(domainsChannel)
	} else if crawl {
		fmt.Println("crawl")
	} else if len(seed) > 0 {
		go generators.GenerateBySeeder(domainsChannel, seed)
	}

	for domain := range domainsChannel {
		var d models.Domain
		models.Db.FirstOrCreate(&d, domain)
	}
}