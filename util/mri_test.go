package util

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadMRIRaw(t *testing.T) {
	pwd, _ := os.Getwd()
	fullpath := pwd + "/../data/" + "mri.raw"
	fmt.Printf("file path:%s\n", fullpath)
	LoadMRIRaw(fullpath, 200, 160, 160)
}
