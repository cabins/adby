package core

import (
	"fmt"
	"testing"
)

func TestGetInstalledVersion(t *testing.T) {
	fmt.Println(GetInstalledVersion("com.vivo.demo"))
}
