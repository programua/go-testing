package dependency_injection

import (
	"bytes"
	"fmt"
)

func Greed(writer *bytes.Buffer, name string) {
	fmt.Printf("Hello, %s", name)
}

