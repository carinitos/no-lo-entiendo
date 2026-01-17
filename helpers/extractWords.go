package helpers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// get words from file
func ListAndReadTextFile() ([]string, error) {
	dir := "input"
	files, err := filepath.Glob(filepath.Join(dir, "*.txt"))
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no text files found in directory: %s", dir)
	}

	fmt.Println("Available files: ")
	for i, file := range files {
		fmt.Printf("%d: %s\n", i+1, filepath.Base(file))
	}

	var choice int
	fmt.Print("Enter the number of the file to be cut up: ")
	_, err = fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(files) {
		return nil, fmt.Errorf("invalid choice")
	}

	selectedFile := files[choice-1]
	var words []string

	file, err := os.Open(selectedFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, strings.Fields(scanner.Text())...)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
