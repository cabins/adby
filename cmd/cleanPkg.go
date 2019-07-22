package cmd

import (
	"adby/core"

	"github.com/spf13/cobra"
)

// cleanPkgCmd represents the cleanPkg command
var cleanPkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "清除应用信息（数据，缓存）",
	Long:  `清除应用信息`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		for _, pkg := range args {
			core.CleanPkg(pkg)
		}
	},
}

func init() {
	cleanCmd.AddCommand(cleanPkgCmd)
}
