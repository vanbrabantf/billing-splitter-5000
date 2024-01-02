package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Names struct {
	Names []Name `json:"names"`
}

type Name struct {
	Name string `json:"name"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory_path> eg: ./incoming")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	err := groupFilesInFolders(dirPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Files grouped successfully!")

}

func groupFilesInFolders(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
	}

	folderName := ""

	for i, f := range files {
		if i == 0 || i%15 == 0 {
			folderName = nameGenerator()
			os.Mkdir(dirPath + "/"+ folderName, 0755)
		}

		err := os.Rename(dirPath + "/" + f.Name(), dirPath + "/" + folderName + "/" + f.Name())

		if err != nil {
			return err
		}
	}

	return nil
}

func nameGenerator() string {
	fileContent, err := os.Open("names.json")

	if err != nil {
		fmt.Println(err)
	}

	defer fileContent.Close()
	byteResult, _ := io.ReadAll(fileContent)
	var names Names
	json.Unmarshal(byteResult, &names)

	return names.Names[rand.Intn(len(names.Names))].Name

}
