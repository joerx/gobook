package fmt_test

import (
	"fmt"
	"testing"
)

// func main() {
// 	fmt.Println(fmt.IsSpace(" "))
// }

func TestIsSpace(t *testing.T) {
	if !fmt.IsSpace(" ") { // not working for stdlib packages since sources are missing?
		t.Error("Not a space?")
	}
}
