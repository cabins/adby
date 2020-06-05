package cmd

import (
	"adby/core"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var pkgCmd = &cobra.Command{
	Use:   "pkg",
	Short: "软件包操作",
	Long:  `软件包操作`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "查看已安装应用，默认为列出全部",
	Long:  `查看已安装应用，默认为列出全部`,
	Run: func(cmd *cobra.Command, args []string) {
		system, _ := cmd.Flags().GetBool("system")
		third, _ := cmd.Flags().GetBool("third")

		if system {
			core.EachPrint(core.ListApps("-s"))
			return
		}

		if third {
			core.EachPrint(core.ListApps("-3"))
			return
		}

		core.EachPrint(core.ListApps(""))
	},
}

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

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "下载应用的安装包",
	Long:  `下载应用的安装包`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("请指定下载的应用包名")
			return
		}

		dest, _ := cmd.Flags().GetString("dest")
		for _, pkgName := range args {
			// Get pkg info
			uri := core.Info(pkgName).ApkUrl
			core.DownloadFileWithProgress(uri, dest)
		}
	},
}

func init() {
	downloadCmd.Flags().StringP("dest", "d", "", "文件保存路径")

	installCmd.Flags().BoolP("local", "l", false, "从本地安装，默认关闭")

	uninstallCmd.Flags().BoolP("all", "a", false, "是否卸载全部应用")
	uninstallCmd.Flags().BoolP("system", "s", false, "是否卸载全部预装应用的更新")
	uninstallCmd.Flags().BoolP("third", "3", false, "是否卸载全部第三方应用")

	updateCmd.Flags().BoolP("all", "a", false, "是否升级全部应用")

	listCmd.Flags().BoolP("system", "s", false, "只列出系统预装应用")
	listCmd.Flags().BoolP("third", "3", false, "只列出用户安装应用")

	infoCmd.Flags().BoolP("table", "t", false, "以表格的形式进行结果输出")

	pkgCmd.AddCommand(installCmd)
	pkgCmd.AddCommand(uninstallCmd)
	pkgCmd.AddCommand(updateCmd)
	pkgCmd.AddCommand(searchCmd)
	pkgCmd.AddCommand(listCmd)
	pkgCmd.AddCommand(infoCmd)
	pkgCmd.AddCommand(downloadCmd)

	rootCmd.AddCommand(pkgCmd)
}
