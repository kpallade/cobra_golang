package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var formatDate string

// timeCmd represents the timezone command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the current time in a given timezone",
	Long: `Get the current time in a given timezone.
This command takes one argument, the timezone you want to get the current time in.
It returns the current time in RFC1123 format.`,

	// start runTime function when "time" cmd is used
	Run: runTime,
}

func runTime(cmd *cobra.Command, args []string) {

	timezone, err := cmd.Flags().GetString("timezone")
	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(err)
	}

	var currentTime string
	if formatDate != "" {
		currentTime = time.Now().In(location).Format(formatDate)
	} else {
		currentTime = time.Now().In(location).Format(time.RFC3339)
	}
	fmt.Println(currentTime)

}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.Flags().Bool("json", false, "JSON format")
	timeCmd.Flags().StringVarP(&formatDate, "date", "d", "", "Date for which to get the time (ex: \"02/01/2006 15:04\", \"02 Jan 06 15:04\")")
	timeCmd.Flags().StringP("timezone", "t", "", "Timezone you are looking for (ex: \"Europe/Paris\")")
	timeCmd.MarkFlagRequired("timezone")
}
