package readfile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile() ([]string, error) {
	fmt.Println("---ReadFile Func---")

	absPath, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	fmt.Printf("AbsPath: %s\n", absPath)

	fName := "rURLs.txt"
	fPath := filepath.Join(absPath, "Files")
	fPath = filepath.Join(fPath, fName)

	fmt.Printf("FilePath: %s\n", fPath)

	f, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()


	fmt.Println("---URLs---")
	urls := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		urls = append(urls, sc.Text())
	}

	return urls, nil
}
