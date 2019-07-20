package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

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
