package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var absPath = getAbsPath()

func getAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println("absPath:", dir)
	return dir
}
