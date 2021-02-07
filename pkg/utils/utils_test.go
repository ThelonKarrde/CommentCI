package utils_test

import (
	"github.com/ThelonKarrde/CommentCI/pkg/utils"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadFileToString(t *testing.T) {
	testString := "test content"
	content := []byte(testString)
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	readResult := utils.ReadFileToString(tmpfile.Name())
	if testString != readResult {
		t.Error("File read error! " + readResult)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestConvertFilesToStrings(t *testing.T) {
	testStrings := [2]string{"test content 1", "test content 2"}
	tmp1, err := ioutil.TempFile("", "test1")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmp1.Name())
	tmp2, err := ioutil.TempFile("", "test2")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmp2.Name())
	if _, err := tmp1.Write([]byte(testStrings[0])); err != nil {
		log.Fatal(err)
	}
	if _, err := tmp2.Write([]byte(testStrings[1])); err != nil {
		log.Fatal(err)
	}
	readResults := utils.ConvertFilesToStrings([]string{tmp1.Name(), tmp2.Name()})
	for i, r := range readResults {
		if r != testStrings[i] {
			t.Error("File convert error! " + r)
		}
	}
	if err := tmp1.Close(); err != nil {
		log.Fatal(err)
	}
	if err := tmp2.Close(); err != nil {
		log.Fatal(err)
	}
}
