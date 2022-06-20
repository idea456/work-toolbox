/*
Copyright © 2022 Adrienne Rio Wongso Atmojo riowongsoatmojo@gmail.com
*/
package remind

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type By string
const (
	TODAY = "today"
	TOMORROW = "tomorrow"
	THIS_WEEK = "this-week"
)

type Reminder struct {
	title string
	message string
	by By
	postedOn time.Time
}

// global variables for arguments
var (
	by string
)


func addReminder() {

}

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new reminder to notify myself later",
	Long: `Add a new reminder to notify myself later.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {

	AddCmd.Flags().StringVarP(&by, "by", "t", "today", "Set the reminder by a certain time period, available options are: 12-HR time period, today, tomorrow, this-week")
}
