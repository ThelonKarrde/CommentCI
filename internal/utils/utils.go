package utils

import (
	"io/ioutil"
	"log"
)

func ReadFileToString(filePath string) string {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return "Unable to read a file: " + filePath
	}
	return string(fileContent)
}

func ConvertFilesToStrings(fileList []string) []string {
	var stringList []string
	for _, p := range fileList {
		stringList = append(stringList, ReadFileToString(p))
	}
	return stringList
}
