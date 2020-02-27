package cmd

import (
	"fmt"
	"github.com/timoisik/web-map/fetchers"
	"github.com/timoisik/web-map/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/spf13/cobra"
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
	// fetchDomainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchDomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fetchDomains() {
	fmt.Println("fetchDomain called")

	domains := models.ReadAll()

	for _, domain := range domains {
		fmt.Printf("Fetch domain %v\n", domain.GetUrl())

		ip, err := fetchers.FetchDomainIp(*domain)

		if err != nil {
			domain.Delete()
		} else {
			domain.Update(bson.D{
				{"$set", bson.D{
					{"ips", ip},
				}},
				{"$currentDate", bson.D{
					{"fetchedat", true},
				}},
			})
		}
	}
}
