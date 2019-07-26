package core

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

// Version - The version info about this app
type Version struct {
	Version     string   `json:"version"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

// PrintAsJSON - Print version info as json
func (version Version) PrintAsJSON() {
	js, err := json.MarshalIndent(version, "  ", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(js))
}

func (version Version) PrintAsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Version", "Author", "Description"})
	table.Append([]string{version.Version, strings.Join(version.Authors, ", "), version.Description})
	table.Render()
}

// Ver - The version info of this app
var Ver Version = Version{
	Version:     "1.6",
	Authors:     []string{"KongLingCun"},
	Description: "优化部分命令，添加更多命令",
}
