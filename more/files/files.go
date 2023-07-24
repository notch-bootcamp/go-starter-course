package main

import (
	"fmt"
	"io"
	"os"
)

func copyWithHeader(filename string, header string) error {
	filePath, err := getAbsolutePath(filename)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	data = []byte(fmt.Sprintf("#%s\n\n\t%s", header, string(data)))
	headerFilePath, err := getAbsolutePath("header-" + filename)
	return os.WriteFile(headerFilePath, data, 0644)
}

func printFile(filename string) error {
	filePath, err := getAbsolutePath(filename)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func getAbsolutePath(filename string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/more/files/%s", pwd, filename), nil
}

func main() {
	err := printFile("example.txt")
	if err != nil {
		fmt.Printf("Error occured while printing a file %s", err)
		return
	}

	err = copyWithHeader("example.txt", "Lorem ipsum")
	if err != nil {
		fmt.Printf("Error occured while doing a copy of a file %s", err)
		return
	}

	err = printFile("header-example.txt")
	if err != nil {
		fmt.Printf("Error occured while printing a header file %s", err)
		return
	}
}
