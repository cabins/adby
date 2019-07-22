package core

import (
	"fmt"
	"testing"
)

func TestNeedToUpdate(t *testing.T) {
	fmt.Println(NeedToUpdate("1.2", "1.0"))
	fmt.Println(NeedToUpdate("1.2.1", "1.2"))
	fmt.Println(NeedToUpdate("1.2", "1.3"))
	fmt.Println(NeedToUpdate("1.2", "1.3.1"))
	fmt.Println(NeedToUpdate("1.2", "1.2..0.1"))
	fmt.Println(NeedToUpdate("1.2", "1.2Beta"))
}
