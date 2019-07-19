package core

import (
	"log"
	"os/exec"
)

// UninstallApp - Uninstall the App
func UninstallApp(pkg string) {
	cmd := exec.Command("adb", "uninstall", pkg)
	cmd.Run()

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}
