package cmd

import (
	"adby/core"
	"log"

	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载应用",
	Long:  `卸载应用`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请指定要卸载的应用")
			return
		}

		for _, pkg := range args {
			core.UninstallApp(pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
