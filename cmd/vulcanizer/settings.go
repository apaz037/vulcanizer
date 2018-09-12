package main

import (
	"fmt"

	"github.com/github/vulcanizer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdSettings)
}

var cmdSettings = &cobra.Command{
	Use:   "settings",
	Short: "Display all the settings of the cluster.",
	Long:  `This command displays all the transient and persisent settings that have been set on the given cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, port := getConfiguration()
		v := vulcanizer.NewClient(host, port)
		rows, headers := v.GetSettings()

		fmt.Println(renderTable(rows, headers))
	},
}
