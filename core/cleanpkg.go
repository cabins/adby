package core

import (
	"log"
	"os/exec"
)

func CleanPkg(pkg string) {
	log.Printf("正在清理软件包%s的数据……\n", pkg)
	cmd := exec.Command("adb", "shell", "pm", "clear", pkg)

	err := cmd.Run()
	if err != nil {
		log.Println("清除数据失败。请重试……")
	}
}
