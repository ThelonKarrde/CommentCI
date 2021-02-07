package main

import (
	"fmt"
	"github.com/ThelonKarrde/CommentCI/pkg/config"
)

func main() {
	data := config.ReadConfig()
	fmt.Println(data)
}
