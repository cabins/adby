package core

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	if verifyApkFile(appInfo, filename) {
		log.Println("MD5校验通过，准备进行安装……")
		InstallFromLocal(filename)
	} else {
		log.Println("MD5校验不通过，跳过安装……")
		return
	}
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

func verifyApkFile(app App, filename string) bool {
	expectMd5 := app.ApkMd5
	actualMd5, _ := getFileMd5(filename)

	log.Printf("预期MD5：%s", expectMd5)
	log.Printf("实际MD5：%s", actualMd5)

	if expectMd5 == actualMd5 {
		return true
	}

	return false
}

func getFileMd5(filename string) (string, error) {
	cachedir := GetCacheFolder()
	filePath := filepath.Join(cachedir, filename)

	var returnMD5String string

	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = strings.ToUpper(hex.EncodeToString(hashInBytes))
	return returnMD5String, nil

}
