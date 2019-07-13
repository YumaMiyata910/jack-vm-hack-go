package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("引数は1つのみ指定してください。 args_count: %d", len(args)-1)
	}

	path := args[1]
	err := readFiles(path)
	if err != nil {
		log.Fatal(err)
	}
}

func readFiles(path string) error {
	var err error

	files, err := ioutil.ReadDir(path)
	if err != nil {
		err = fmt.Errorf("指定されたファイルまたはディレクトリが存在しません。 path: %s", path)
	}

	for _, file := range files {
		if file.IsDir() {
			err = readFiles(filepath.Join(path, file.Name()))
			continue
		}

		if !strings.HasSuffix(file.Name(), ".vm") {
			continue
		}

		err = convert(path)
	}

	return err
}

func convert(path string) error {
	var err error
	readfile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer readfile.Close()

	// sc := bufio.NewScanner(readfile)
	// p := parser.NewParser(sc)

	fmt.Println(readfile)

	return err
}
