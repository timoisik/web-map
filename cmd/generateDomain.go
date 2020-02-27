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
		generateDomains()
	},
}

func init() {
	generateCmd.AddCommand(generateDomainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateDomainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateDomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateDomains() {
	fmt.Println("generateDomain called")

	domainsChannel := make(chan models.Domain)
	go generators.GenerateDomains(domainsChannel)

	for domain := range domainsChannel {
		domain.Create()
	}
}