/*
Copyright © 2022 Adrienne Rio Wongso Atmojo riowongsoatmojo@gmail.com
*/

package notify

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	notifyType string
)

// notifyMeCmd represents the notifyMe command
var NotifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "notify pings you about pull request reviews and comments, build progress, meetings or your own reminders.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Logic Here
		Authenticate()
	},
}

func init() {
	NotifyCmd.Flags().StringVarP(&notifyType, "type", "t", "", "The type of ping to notify");
	
	if err := NotifyCmd.MarkFlagRequired("type"); err != nil {
		fmt.Println(err)
	}

}
