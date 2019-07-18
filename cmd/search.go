package cmd

import (
	"adby/core"
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "通过应用关键字来搜索应用",
	Long:  `通过应用关键字来搜索应用`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("请输入一个关键字")
			return
		}

		name := args[0]
		core.SearchPkg(name).PrintAsTable()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
