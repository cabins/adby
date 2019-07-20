package core

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

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
