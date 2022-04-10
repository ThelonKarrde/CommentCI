package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadFileToString(filePath string) string {
	fileContent, err := os.ReadFile(filePath)
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
		fmt.Println(stringList)
	}
	return stringList
}
