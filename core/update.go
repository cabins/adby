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

	return strings.TrimSpace(string(verList[1]))

}

func GetRepoVersion(pkg string) string {
	return Info(pkg).VersionName
}

func NeedToUpdate(localVersion, repoVersion string) bool {
	// 版本号对比，来判断是否需要升级
	lVer := strings.Split(localVersion, ".")
	rVer := strings.Split(repoVersion, ".")

	length := len(lVer)
	if len(rVer) < length {
		length = len(rVer)
	}

	for i := 0; i < length; i++ {
		if rVer[i] > lVer[i] {
			return true
		}
	}

	if len(rVer) > len(lVer) {
		return true
	}

	return false
}

func UpdateApp(pkg string) {
	localVersion := GetInstalledVersion(pkg)
	repoVersion := GetRepoVersion(pkg)

	if NeedToUpdate(localVersion, repoVersion) {
		log.Printf("%s已安装版本%s，发现新版本%s，开始升级……\n", pkg, localVersion, repoVersion)
		InstallFromNet(pkg)
	} else {
		log.Printf("%s最新版本为%s，已安装版本%s……\n", pkg, repoVersion, localVersion)
	}
}
