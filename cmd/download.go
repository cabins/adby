package cmd

import (
	"adby/core"
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
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
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("dest", "d", "", "文件保存路径")
}
