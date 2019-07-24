package cmd

import (
	"adby/core"
	"log"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "操作Android常用设置",
	Long:  `操作Android常用设置`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var setwifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "开启/关闭Wi-Fi开关",
	Long:  `开启/关闭Wi-Fi开关`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Println("请指定操作：on/off")
			return
		}

		if args[0] == "on" {
			core.TurnOnWifi()
			return
		}

		if args[0] == "off" {
			core.TurnOffWifi()
		}
	},
}

func init() {
	setCmd.AddCommand(setwifiCmd)
	rootCmd.AddCommand(setCmd)
}
