package core

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetInstalledVersion(pkg string) string {
	str := fmt.Sprintf("dumpsys package %s | grep versionName", pkg)
	ver, err := exec.Command("adb", "shell", str).Output()
	if err != nil {
		log.Println(err)
		return ""
	}

	verList := strings.Split(string(ver), "=")
	if len(verList) != 2 {
		return ""
	}

	return string(verList[1])

}

func GetRepoVersion(pkg string) string {
	return Info(pkg).VersionName
}

func NeedToUpdate(localVersion, repoVersion string) bool {
	return false
}

func TransVersion(version string) []string {
	return strings.Split(version, ".")
}
