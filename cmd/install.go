package cmd

import (
	"adby/core"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装软件包",
	Long:  `安装软件包`,
	Run: func(cmd *cobra.Command, args []string) {
		isLocal, _ := cmd.Flags().GetBool("local")

		// install from network
		for _, pkg := range args {
			if isLocal {
				core.InstallFromLocal(pkg)
				continue
			}

			core.InstallFromNet(pkg)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolP("local", "l", false, "从本地安装，默认关闭")
}
