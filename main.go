package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/yookoala/realpath"
	"github.com/truestblue/php-transformer/parser"
	"github.com/truestblue/php-transformer/php5"
	"github.com/truestblue/php-transformer/php7"
	printer2 "github.com/truestblue/php-transformer/printer"
	"github.com/truestblue/php-transformer/visitor"
	"io"
	"strings"
	"github.com/truestblue/php-transformer/dictionary"
)

const (
	DIR = iota
	CPY = iota
	SCR = iota
)

var wg sync.WaitGroup
var usePhp5 *bool
var noDump *bool

var base = ""
var outFilePath = ""

func main() {
	usePhp5 = flag.Bool("php5", false, "use PHP5 parserWorker")
	noDump = flag.Bool("noDump", false, "disable dumping to stdout")
	flag.Parse()

	pathCh := make(chan string)
	resultCh := make(chan parser.Parser)

	dictionary.InitMapping()

	// run 4 concurrent parserWorkers
	for i := 0; i < 4; i++ {
		go parserWorker(pathCh, resultCh)
	}

	// run printer goroutine
	go printer(resultCh)

	// process files
	processPath(flag.Args(), pathCh)

	// wait the all files done
	wg.Wait()
	close(pathCh)
	close(resultCh)
}

func processPath(pathList []string, pathCh chan<- string) {
	for _, pathCur := range pathList {
		real, err := realpath.Realpath(pathCur)
		checkErr(err)

		fmt.Println(real)

		base = real + "/"
		outFilePath = pathCur + "-ps"
		os.MkdirAll(outFilePath, os.ModePerm)

		err = filepath.Walk(real, func(pathCur string, f os.FileInfo, err error) error {
			if !f.IsDir() && filepath.Ext(pathCur) == ".php" { //TODO: .php is not the only acceptable file extension.
				wg.Add(1)
				pathCh <- pathCur
			} else if f.IsDir() {
				printOut(pathCur, DIR)
			} else {
				printOut(pathCur, CPY)
			}
			return nil
		})
		checkErr(err)
	}
}

func parserWorker(pathCh <-chan string, result chan<- parser.Parser) {
	var parserWorker parser.Parser

	for {
		path, ok := <-pathCh
		if !ok {
			return
		}

		src, _ := os.Open(path)

		if *usePhp5 {
			parserWorker = php5.NewParser(src, path)
		} else {
			parserWorker = php7.NewParser(src, path)
		}

		parserWorker.Parse()
		result <- parserWorker
	}
}

func printer(result <-chan parser.Parser) {
	for {
		parserWorker, ok := <-result
		if !ok {
			return
		}

		fmt.Printf("==> %s\n", parserWorker.GetPath())

		for _, e := range parserWorker.GetErrors() {
			fmt.Println(e)
		}

		if !*noDump {
			nsResolver := visitor.NewNamespaceResolver()
			parserWorker.GetRootNode().Walk(nsResolver)

			dumper := visitor.Dumper{
				Writer:     os.Stdout,
				Indent:     "  | ",
				Comments:   parserWorker.GetComments(),
				Positions:  parserWorker.GetPositions(),
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		}

		fileOut := printOut(parserWorker.GetPath(), SCR)

		file, err := os.Create(fileOut)
		checkErr(err)

		p := printer2.NewPrinter(file, "	")
		p.Print(parserWorker.GetRootNode())
		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printOut(pathCur string, printType int) string {
	sf := strings.SplitAfter(pathCur, base)
	sf[0] = outFilePath
	fileOut := strings.Join(sf, "/")

	switch printType {
	case DIR:
		os.MkdirAll(fileOut, os.ModePerm)
	case CPY:
		copyFile(pathCur, fileOut)
	case SCR:
		//just pass back file name.
	default:
		//panic

	}
	return fileOut

}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
