package cmd

import (
	"adby/core"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "执行清理动作",
	Long:  `执行清理动作`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "清理软件包缓存",
	Long:  `清理软件包缓存`,
	Run: func(cmd *cobra.Command, args []string) {
		cacheFolder := core.GetCacheFolder()

		// Delete all the apk files in cacheFolder
		fi, err := ioutil.ReadDir(cacheFolder)
		if err != nil {
			log.Println(err)
			return
		}

		for _, f := range fi {
			if strings.HasSuffix(f.Name(), "apk") {
				log.Println("正在删除：", f.Name())
				os.RemoveAll(filepath.Join(cacheFolder, f.Name()))
			}
		}
	},
}

func init() {
	cleanCmd.AddCommand(cacheCmd)
	rootCmd.AddCommand(cleanCmd)
}
