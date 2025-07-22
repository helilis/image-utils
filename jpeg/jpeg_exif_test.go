package jpeg

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func Test_addExifInfo(t *testing.T) {
	p, err := filepath.Abs("../data/01.jpg")
	if err != nil {
		t.Fatal("get file fail, err:" + err.Error())
	}
	//
	f, err := os.ReadFile(p)
	if nil != err {
		t.Fatal("read file fail, err:" + err.Error())
	}
	b, err := addExifInfo(bytes.NewReader(f), "exif_test", "exif_value")
	if nil != err {
		t.Fatal("add exif info fail, err:" + err.Error())
	}
	output, _ := filepath.Abs("../data/01_exif.jpg")
	err = os.WriteFile(output, b, 0644)
}
