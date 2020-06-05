package core

import (
	"log"
	"os/exec"
)

// StartApp - Launch up an app by package name
func StartApp(pkgname string) {
	exec.Command("adb", "shell", "monkey", "-p", pkgname, "-c", "android.intent.category.LAUNCHER", "1").Run()
}

// StopApp - Force stop an app by package name
func StopApp(pkgname string) {
	exec.Command("adb", "shell", "am", "force-stop", pkgname).Run()
}

// CleanApp - Clean an app's data by package name
func CleanApp(pkg string) {
	log.Printf("正在清理软件包%s的数据……\n", pkg)
	cmd := exec.Command("adb", "shell", "pm", "clear", pkg)

	err := cmd.Run()
	if err != nil {
		log.Println("清除数据失败。请重试……")
	}
}
