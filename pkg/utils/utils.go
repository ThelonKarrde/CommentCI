package utils

import (
	"io/ioutil"
	"log"
)

func fileToString(filePath string) string {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return "Unable to read a file: " + filePath
	}
	return string(fileContent)
}

func FilesToStrings(fileList []string) []string {
	var textList []string
	for _, p := range fileList {
		textList = append(textList, fileToString(p))
	}
	return textList
}
