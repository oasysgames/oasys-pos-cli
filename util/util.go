package util

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}
