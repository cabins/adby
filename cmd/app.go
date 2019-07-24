package cmd

import (
	"adby/core"
	"log"

	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "App操作，启动/停止/清除数据",
	Long:  `App操作，启动/停止/清除数据`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var appcleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清除应用的数据",
	Long:  `清除应用的数据`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请指定清除数据的软件包名")
			return
		}

		for _, pkg := range args {
			core.CleanPkg(pkg)
		}
	},
}

var appstartCmd = &cobra.Command{
	Use:   "start",
	Short: "通过包名启动应用",
	Long:  `通过包名启动应用`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请指定启动应用的包名")
			return
		}

		for _, pkg := range args {
			core.AppStart(pkg)
		}
	},
}

var appstopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止应用的运行",
	Long:  `停止应用的运行`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请指定要关闭的应用包名")
			return
		}

		for _, pkg := range args {
			core.AppStop(pkg)
		}
	},
}

func init() {
	appCmd.AddCommand(appcleanCmd)
	appCmd.AddCommand(appstartCmd)
	appCmd.AddCommand(appstopCmd)
	rootCmd.AddCommand(appCmd)
}
