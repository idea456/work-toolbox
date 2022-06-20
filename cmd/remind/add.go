/*
Copyright © 2022 Adrienne Rio Wongso Atmojo riowongsoatmojo@gmail.com
*/
package remind

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/spf13/cobra"
)

type By string

const (
	TODAY     = "today"
	TOMORROW  = "tomorrow"
	THIS_WEEK = "this-week"
)

type Reminder struct {
	Title    string
	Message  string
	By       By
	PostedOn time.Time
}

// global variables for arguments
var (
	by string
)

func addReminder() {

}

func remind(title string, message string, by By) {
	reminder := Reminder{
		Title:    title,
		Message:  message,
		By:       by,
		PostedOn: time.Now(),
	}

	bytes, _ := json.Marshal(reminder)
	fmt.Println(string(bytes))
}

func openSchedule() *os.File {
	file, err := os.OpenFile("schedule.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to read schedules: ^%v", err)
	}
	return file
}

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new reminder to notify myself later",
	Long:  `Add a new reminder to notify myself later.`,
	Run: func(cmd *cobra.Command, args []string) {
		reminder := Reminder{
			Title:    "Code review",
			Message:  "this link",
			By:       TODAY,
			PostedOn: time.Now(),
		}

		file := openSchedule()
		r, _ := json.MarshalIndent(reminder, "", "")

		// s := gocron.NewScheduler(time.UTC)

		// s.Every(2).Seconds().Do(func() {
		// 	remind("Code Review", "this link", TODAY)
		// })

		defer file.Close()

		if _, err := file.WriteString(string(r)); err != nil {
			panic(err)
		}

		// s.StartBlocking()
	},
}

func init() {

	AddCmd.Flags().StringVarP(&by, "by", "t", "today", "Set the reminder by HH:MM, available options are: 12-HR time period, today, tomorrow, this-week")
}
