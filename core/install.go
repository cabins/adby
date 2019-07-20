package core

import (
	"log"
	"os/exec"
	"path/filepath"
)

// InstallFromNet - Install From Network
func InstallFromNet(pkg string) {
	appInfo := Info(pkg)
	if appInfo.ApkUrl == "" {
		log.Printf("没有找到%s的相关信息，跳过安装……\n", pkg)
		return
	}

	// 2, Download the APK
	log.Println("下载安装包……")
	DownloadFileWithProgress(appInfo.ApkUrl, "")
	filename := appInfo.GetFileName()

	// 3, Install the APK
	log.Println("准备安装……")
	InstallFromLocal(filename)
}

// InstallFromLocal - Install From Local
func InstallFromLocal(filename string) {
	cachedir := GetCacheFolder()

	fileName := filepath.Join(cachedir, filename)

	cmd := exec.Command("adb", "install", "-d", "-r", fileName)
	cmd.Run()
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("安装完成")
}
