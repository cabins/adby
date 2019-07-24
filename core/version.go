package core

import (
	"encoding/json"
	"fmt"
)

type Version struct {
	Version     string   `json:"version"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
}

func (self Version) PrintAsJson() {
	js, err := json.MarshalIndent(self, "  ", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(js))
}

var Ver Version = Version{
	Version:     "1.5",
	Authors:     []string{"KongLingCun"},
	Description: "优化部分命令，添加更多命令",
}
