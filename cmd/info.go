package cmd

import (
	"adby/core"
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看应用的信息",
	Long:  `查看应用的信息`,
	Run: func(cmd *cobra.Command, args []string) {
		table, _ := cmd.Flags().GetBool("table")

		if len(args) == 0 {
			fmt.Println("请输入一个应用的包名")
			return
		}
		if table {
			core.Info(args[0]).PrintAsTable()
			return
		}
		core.Info(args[0]).PrintAsJson()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().BoolP("table", "t", false, "以表格的形式进行结果输出")
}
