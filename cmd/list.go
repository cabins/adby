package cmd

import (
	"adby/core"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "查看已安装应用，默认为列出全部",
	Long:  `查看已安装应用，默认为列出全部`,
	Run: func(cmd *cobra.Command, args []string) {
		system, _ := cmd.Flags().GetBool("system")
		third, _ := cmd.Flags().GetBool("third")

		if system {
			fmt.Println(core.ListApps("-s"))
			return
		}

		if third {
			fmt.Println(core.ListApps("-3"))
			return
		}

		fmt.Println(core.ListApps(""))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("system", "s", false, "只列出系统预装应用")
	listCmd.Flags().BoolP("third", "3", false, "只列出用户安装应用")
}
