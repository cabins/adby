package core

import (
	"fmt"
	"log"
	"os/exec"
)

// Doctor - Check if adb is installed on PATH
func Doctor() {
	_, err := exec.LookPath("adb")
	if err != nil {
		log.Println("Not found the 'adb' from your PATH.")
		return
	}

	if out, err := exec.Command("adb", "version").Output(); err != nil {
		log.Println(err)
		return
	} else {
		fmt.Printf("\nCongratulations! ADB is found!\n\n")
		fmt.Println("-----------------------------------------------------")
		fmt.Println(string(out))
	}
}
