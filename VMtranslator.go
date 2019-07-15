package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/YumaMiyata910/jack-vm-hack-go/parser"
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

	sc := bufio.NewScanner(readfile)
	p := parser.NewParser(sc)

	filename := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	fPath := filepath.Dir(path) + filename
	writefile, err := os.Create(fPath + ".asm")
	if err != nil {
		return fmt.Errorf("新規ファイルを作成できません。 %v", err)
	}
	defer writefile.Close()

	for p.HasMoreCommands() {
		if err = p.ScannerError(); err != nil {
			return fmt.Errorf("ファイルの読み込みに失敗しました。path:【%s】", path)
		}

		p.Advance()
		if p.Text() == "" {
			continue
		}

	}

	return err
}
