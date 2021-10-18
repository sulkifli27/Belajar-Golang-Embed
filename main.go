package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed fotoku.jpg
var fotoku []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("fotoku_terbaru.jpg", fotoku, fs.ModePerm )

	if err != nil {
		panic(err)
	}

	dirEntries, _	:= path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir(){
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}