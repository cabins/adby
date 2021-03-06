package cmd

import (
	"adby/core"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "操作Android常用设置",
	Long:  `操作Android常用设置`,
	Run: func(cmd *cobra.Command, args []string) {
		deviceInfoCmd.Run(deviceInfoCmd, args)
	},
}

var wifiCmd = &cobra.Command{
	Use:   "wifi",
	Short: "开启/关闭Wi-Fi开关",
	Long:  `开启/关闭Wi-Fi开关`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Println("请指定操作：on/off/prefer")
			return
		}

		if args[0] == "on" {
			core.WifiOn()
			return
		}

		if args[0] == "off" {
			core.WifiOff()
			return
		}

		if args[0] == "prefer" {
			core.WifiPrefer()
			return
		}
	},
}

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "开启/关闭数据网络开关",
	Long:  `开启/关闭数据网络开关`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Println("请指定操作：on/off/prefer")
			return
		}

		if args[0] == "on" {
			core.DataOn()
			return
		}

		if args[0] == "off" {
			core.DataOff()
			return
		}

		if args[0] == "prefer" {
			core.DataPrefer()
			return
		}
	},
}

var statusBarCmd = &cobra.Command{
	Use:     "statusbar",
	Aliases: []string{"sb"},
	Short:   "操作通知栏和快捷中心",
	Long:    `操作通知栏和快捷中心`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Println("请指定操作：1 expand/0 collapse/2 settings/3 forbid/ 4 resume")
			return
		}

		if args[0] == "0" || args[0] == "collapse" {
			core.StatusBarCollapse()
			return
		}

		if args[0] == "1" || args[0] == "expand" {
			core.StatusBarExpand()
			return
		}

		if args[0] == "2" || args[0] == "settings" {
			core.StatusBarSettings()
			return
		}

		if args[0] == "3" || args[0] == "forbid" {
			core.StatusBarForbid()
			return
		}

		if args[0] == "4" || args[0] == "resume" {
			core.StatusBarResume()
			return
		}
	},
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Android时间操作",
	Long:  `Android时间操作`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := cmd.Flags().GetInt("year")
		month, _ := cmd.Flags().GetInt("month")
		week, _ := cmd.Flags().GetInt("week")
		day, _ := cmd.Flags().GetInt("day")
		hour, _ := cmd.Flags().GetInt("hour")
		minute, _ := cmd.Flags().GetInt("minute")
		auto, _ := cmd.Flags().GetBool("auto")

		if auto {
			core.DeviceAutoTime()
			return
		}

		var needSet bool
		for _, item := range []int{year, month, week, day, hour, minute} {
			if item != 0 {
				needSet = true
			}
		}

		if needSet {
			core.DeviceTimeChange(year, month, week, day, hour, minute)
		}
		fmt.Println(core.DeviceCurrentTime())
	},
}

var deviceInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "获取设备信息",
	Long:  `获取设备信息`,
	Run: func(cmd *cobra.Command, args []string) {
		table, _ := cmd.Flags().GetBool("table")
		if table {
			core.GetDeviceInfo().PrintAsTable()
			return
		}
		core.GetDeviceInfo().PrintAsJSON()
	},
}

var screenCmd = &cobra.Command{
	Use:   "screen",
	Short: "屏幕相关操作",
	Long:  "屏幕相关操作",
	Run: func(cmd *cobra.Command, args []string) {
		brightness, _ := cmd.Flags().GetInt("brightness")
		autoBrightness, _ := cmd.Flags().GetInt("auto")
		timeout, _ := cmd.Flags().GetInt("timeout")

		if autoBrightness == 1 {
			log.Println(core.ScreenBrightMode(1))
			return
		} else {
			log.Println(core.ScreenBrightMode(0))
		}

		if brightness != -1 {
			log.Println(core.ScreenBrightness(strconv.Itoa(brightness)))
			return
		} else {
			log.Println(core.ScreenBrightness(""))
		}

		if timeout != -1 {
			log.Println(core.ScreenOffTimeout(timeout))
			return
		} else {
			log.Println(core.ScreenOffTimeout(-1))
		}
		core.GetScreenInfo().PrintAsTable()
	},
}

func init() {
	deviceCmd.AddCommand(wifiCmd)
	deviceCmd.AddCommand(dataCmd)
	deviceCmd.AddCommand(statusBarCmd)

	timeCmd.Flags().IntP("minute", "M", 0, "需要增减的分钟数")
	timeCmd.Flags().IntP("hour", "H", 0, "需要增减的小时数")
	timeCmd.Flags().IntP("day", "d", 0, "需要增减的天数")
	timeCmd.Flags().IntP("week", "w", 0, "需要增减的周数")
	timeCmd.Flags().IntP("month", "m", 0, "需要增减的月数")
	timeCmd.Flags().IntP("year", "y", 0, "需要增减的年数")
	timeCmd.Flags().BoolP("auto", "a", false, "自动时间设定")
	deviceCmd.AddCommand(timeCmd)

	deviceInfoCmd.Flags().BoolP("table", "t", false, "是否以表格格式打印")
	deviceCmd.AddCommand(deviceInfoCmd)

	screenCmd.Flags().IntP("brightness", "b", -1, "设置屏幕亮度，默认为100")
	screenCmd.Flags().IntP("auto", "a", -1, "设置屏幕自动亮度, 1自动，-1不自动")
	screenCmd.Flags().IntP("timeout", "t", -1, "设置自动息屏时间，单位毫秒，默认为30秒")
	deviceCmd.AddCommand(screenCmd)

	rootCmd.AddCommand(deviceCmd)
}
