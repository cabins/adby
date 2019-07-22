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
		all, _ := cmd.Flags().GetBool("all")
		system, _ := cmd.Flags().GetBool("system")
		third, _ := cmd.Flags().GetBool("third")

		if all {
			for _, pkg := range core.ListApps("") {
				core.UninstallApp(pkg)
			}
			return
		}

		if system {
			for _, pkg := range core.ListApps("-s") {
				core.UninstallApp(pkg)
			}
			return
		}

		if third {
			for _, pkg := range core.ListApps("-3") {
				core.UninstallApp(pkg)
			}
			return
		}

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
	uninstallCmd.Flags().BoolP("all", "a", false, "是否卸载全部应用")
	uninstallCmd.Flags().BoolP("system", "s", false, "是否卸载全部预装应用的更新")
	uninstallCmd.Flags().BoolP("third", "3", false, "是否卸载全部第三方应用")
}
