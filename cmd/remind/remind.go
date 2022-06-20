/*
Copyright © 2022 Adrienne Rio Wongso Atmojo riowongsoatmojo@gmail.com
*/

package remind

import (
	"github.com/spf13/cobra"
)

// remindCmd represents the remind command
var RemindCmd = &cobra.Command{
	Use:   "remind",
	Short: "remind command reminds you of tasks you saved using this command.",
	Long: `remind commands is a reminder tool to help remind me of tasks I actually need to remember cuz I can't always use sticky notes or remember them always...
	`,
	Example: `remind add "Fix that freaking tooltip later after finishing block user feature!" --by this-week`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	RemindCmd.AddCommand(AddCmd);
}
