package core

import (
	"log"
	"os/exec"
)

func StartApp(pkgname string) {
	exec.Command("adb", "shell", "monkey", "-p", pkgname, "-c", "android.intent.category.LAUNCHER", "1").Run()
}

func StopApp(pkgname string) {
	exec.Command("adb", "shell", "am", "force-stop", pkgname).Run()
}

func CleanApp(pkg string) {
	log.Printf("正在清理软件包%s的数据……\n", pkg)
	cmd := exec.Command("adb", "shell", "pm", "clear", pkg)

	err := cmd.Run()
	if err != nil {
		log.Println("清除数据失败。请重试……")
	}
}
