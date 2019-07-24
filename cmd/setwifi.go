/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"adby/core"
	"log"

	"github.com/spf13/cobra"
)

// setwifiCmd represents the setwifi command
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
}
