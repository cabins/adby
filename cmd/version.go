package cmd

import (
	"adby/core"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看工具的版本号",
	Long:  `查看工具的版本号`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Ver.PrintAsJson()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
