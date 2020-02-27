package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/timoisik/web-map/fetchers"
	"github.com/timoisik/web-map/models"
)

// fetchDomainCmd represents the fetchDomain command
var fetchDomainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Fetches information about an domain.",
	Long: `Fetches information like IP addresses about an domain.`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchDomains()
	},
}

func init() {
	fetchCmd.AddCommand(fetchDomainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchDomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchDomains() {
	fmt.Println("fetchDomain called")

	var domains []models.Domain
	models.Db.Find(&domains)

	for _, domain := range domains {
		fmt.Printf("Fetch domain %v\n", domain.GetUrl())

		ips, err := fetchers.FetchDomainIp(domain)
		if err == nil {
			for _, ipAddress := range ips {

				// Check if ip address exists in db
				var ip models.Ip
				models.Db.FirstOrCreate(&ip, models.Ip{Address: ipAddress.String()})
				models.Db.Model(domain).Association("IpAddresses").Append(ip)
			}
		}
	}
}
