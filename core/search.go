package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Package struct {
	Obj Object `json:"obj"`
}

type Object struct {
	AppDetail []App `json:"appDetails"`
}

type App struct {
	FileSize    int    `json:"fileSize"`
	ApkMd5      string `json:"apkMd5"`
	ApkUrl      string `json:"apkUrl"`
	AppName     string `json:"appName"`
	AuthorName  string `json:"authorName"`
	IconUrl     string `json:"iconUrl"`
	PkgName     string `json:"pkgName"`
	VersionName string `json:"versionName"`
}

func (self Object) PrintAsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名称", "包名", "文件大小", "下载地址", "版本"})

	for _, v := range self.AppDetail {
		table.Append([]string{v.AppName, v.PkgName, fmt.Sprintf("%s", Byte2Human(v.FileSize)), strings.Split(v.ApkUrl, "?")[0], v.VersionName})
	}

	table.Render()
}

func (self App) PrintAsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"名称", "包名", "文件大小", "下载地址", "版本"})

	table.Append([]string{self.AppName, self.PkgName, fmt.Sprintf("%s", Byte2Human(self.FileSize)), strings.Split(self.ApkUrl, "?")[0], self.VersionName})

	table.Render()
}

func (self App) GetFileName() string {
	urlValue, err := url.Parse(self.ApkUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	destName := urlValue.Query().Get("fsname")
	return destName
}

func (self App) PrintAsJson() {
	res, err := json.MarshalIndent(self, "  ", "  ")
	if err != nil {
		fmt.Println("序列化App信息出错")
		return
	}
	fmt.Println(string(res))
}

func SearchPkg(query string) Object {
	url := fmt.Sprintf("https://sj.qq.com/myapp/searchAjax.htm?kw=%s", query)

	resp, err := http.Get(url)
	if err != nil {
		return Object{}
	}

	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Object{}
	}

	pkgInfo := Package{}
	err = json.Unmarshal(res, &pkgInfo)
	if err != nil {
		return Object{}
	}

	return pkgInfo.Obj
}
