package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

func DownloadFileWithProgress(uri string, destPath string) {
	if destPath == "" {
		destPath = GetCacheFolder()
	}

	// Make dir if destpath does not exist and not be empty
	if destPath != "" {
		if _, err := os.Stat(destPath); os.IsNotExist(err) {
			e := os.MkdirAll(destPath, os.ModePerm)
			if e != nil {
				fmt.Println(err)
				destPath = ""
			}
		}
	}

	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Server return non-200 status: %s\n", resp.Status)
		return
	}

	size := resp.ContentLength

	// create dest
	urlValue, err := url.Parse(uri)
	if err != nil {
		fmt.Println(err)
		return
	}

	destName := urlValue.Query().Get("fsname")

	dest, err := os.Create(filepath.Join(destPath, destName))
	if err != nil {
		fmt.Printf("Can't create %s: %v\n", destName, err)
		return
	}
	defer dest.Close()

	p := mpb.New(
		mpb.WithWidth(60),
		mpb.WithRefreshRate(180*time.Millisecond),
	)

	bar := p.AddBar(size, mpb.BarStyle("[=>-|"),
		mpb.PrependDecorators(
			decor.CountersKibiByte("% 6.1f / % 6.1f"),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_MMSS, float64(size)/2048),
			decor.Name(" ] "),
			decor.AverageSpeed(decor.UnitKiB, "% .2f"),
		),
	)

	// create proxy reader
	reader := bar.ProxyReader(resp.Body)

	// and copy from reader, ignoring errors
	io.Copy(dest, reader)

	p.Wait()
}

func Info(pkgName string) App {
	appName := GetAppNameByPkg(pkgName)
	if appName == "" {
		return App{}
	}

	object := SearchPkg(appName)

	for _, v := range object.AppDetail {
		if v.PkgName == pkgName {
			return v
		}
	}

	return App{}
}

func GetAppNameByPkg(pkgName string) string {
	res, err := http.Get(fmt.Sprintf("https://sj.qq.com/myapp/detail.htm?apkName=%s", pkgName))
	if err != nil {
		log.Println(err)
		return ""
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("解析文件详情失败")
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panicln("解析应用详情失败")
		return ""
	}

	appName := doc.Find(".det-main-container .det-ins-data .det-name .det-name-int").Text()

	return appName
}

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

// UninstallApp - Uninstall the App
func UninstallApp(pkg string) {
	log.Printf("正在尝试卸载%s\n", pkg)
	cmd := exec.Command("adb", "uninstall", pkg)
	cmd.Run()

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

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

func ListApps(arg string) []string {
	res, err := exec.Command("adb", "shell", "pm", "list", "packages", arg).Output()
	if err != nil {
		log.Println(err)
		return []string{}
	}

	// fmt.Println(string(res))

	var resList []string
	for _, line := range strings.Split(string(res), "\n") {
		if strings.HasPrefix(line, "package:") {
			resList = append(resList, strings.TrimSpace(strings.Split(line, ":")[1]))
		}
	}
	return resList
}
