package cmd

import (
	"adby/core"
	"log"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "更新应用",
	Long:    `更新应用`,
	Run: func(cmd *cobra.Command, args []string) {
		reArgs := args
		all, _ := cmd.Flags().GetBool("all")
		if all {
			// 升级全部应用
			reArgs = core.ListApps("")
		} else if len(args) == 0 {
			log.Println("请指定应用名称，或使用--all来升级全部")
			return
		}

		for _, pkg := range reArgs {
			core.UpdateApp(pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().BoolP("all", "a", false, "是否升级全部应用")
}
