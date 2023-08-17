package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Printf("Error reading stdin: %s\n", err)
			os.Exit(1)
		}
		return
	}

	for _, fileName := range os.Args[1:] {
		if fileName == "-" {
			_, err := io.Copy(os.Stdout, os.Stdin)
			if err != nil {
				fmt.Printf("Error reading stdin: %s\n", err)
				os.Exit(1)
			}
			continue
		}

		files, err := filepath.Glob(fileName)
		if err != nil {
			fmt.Printf("Error globbing %s: %s\n", fileName, err)
			os.Exit(1)
		}

		for _, file := range files {
			err := catFile(file)
			if err != nil {
				fmt.Printf("Error reading %s: %s\n", file, err)
				os.Exit(1)
			}
		}
	}
}

func catFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return err
	}

	return nil
}
