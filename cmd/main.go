package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

func main() {
	help := flag.Bool("h", false, "Show help")
	helpLong := flag.Bool("help", false, "Show help")
	path := flag.String("p", "", "Path to the directory")
	pathLong := flag.String("path", "", "Path to the directory")
	includeHidden := flag.Bool("i", false, "Include hidden folders")
	includeHiddenLong := flag.Bool("include-hidden", false, "Include hidden folders")

	flag.Parse()

	if *help || *helpLong {
		showHelp()
		return
	}

	dirPath := *path
	if *pathLong != "" {
		dirPath = *pathLong
	}

	if dirPath == "" {
		fmt.Println("Error: Path is required")
		showHelp()
		return
	}

	includeHiddenFolders := *includeHidden || *includeHiddenLong

	err := printDirStructure(dirPath, includeHiddenFolders)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func showHelp() {
	fmt.Printf("Usage: TreeScribe [options]\n")
	fmt.Printf("Options:\n")
	fmt.Printf("  -h, --help             Show help\n")
	fmt.Printf("  -p, --path             Path to the directory\n")
	fmt.Printf("  -i, --include-hidden   Include hidden folders\n")
}

func printDirStructure(root string, includeHidden bool) error {
	fileInfo, err := os.Stat(root)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", root)
	}

	fmt.Println(fileInfo.Name() + "/")
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, _ := filepath.Rel(root, path)
		if relativePath == "." {
			return nil
		}

		if !includeHidden && strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		parts := strings.Split(relativePath, string(filepath.Separator))
		for i := 1; i < len(parts); i++ {
			if i == len(parts)-1 {
				if info.IsDir() {
					fmt.Print("├── ")
				} else {
					fmt.Print("└── ")
				}
			} else {
				fmt.Print("│   ")
			}
		}

		if info.IsDir() {
			fmt.Printf("%s/\n", info.Name())
		} else {
			fmt.Printf("%s\n", info.Name())
		}
		return nil
	})
}
