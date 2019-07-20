package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "更新应用",
	Long:    `更新应用`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
