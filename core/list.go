package core

import (
	"log"
	"os/exec"
	"strings"
)

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
			resList = append(resList, strings.Split(line, ":")[1])
		}
	}
	return resList
}
