package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	switch flag.NArg() {
	case 0:
		s, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if s.Mode()&os.ModeNamedPipe == 0 {
			// no data
			usage()
		}
		d, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		process(d)
	case 1:
		d, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		process(d)
	default:
		usage()
	}
}

func process(data []byte) {
	if len(data) == 0 {
		usage()
	}
	// it already may be a windows file
	data = bytes.ReplaceAll(data, []byte{'\r', '\n'}, []byte{'\n'})
	data = bytes.ReplaceAll(data, []byte{'\n'}, []byte{'\r', '\n'})
	fmt.Print(string(data))
}

func usage() {
	exit("crlfs [filepath] | [stdin]")
}

func exit(format string, a ...interface{}) {
	if len(a) > 0 {
		fmt.Println(format, a)
	} else {
		fmt.Println(format)
	}
	os.Exit(1)
}
