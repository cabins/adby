package core

import (
	"os/exec"
)

func TurnOnWifi() {
	wifiService("enable")
}

func TurnOffWifi() {
	wifiService("disable")
}

func wifiService(operate string) {
	cmd := exec.Command("adb", "shell", "svc", "wifi", operate)
	cmd.Run()
}
