package main

import (
	"flag"
	"github.com/sanemat/go-xgoinstall"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var data []byte
	var fi os.FileInfo
	var err error
	log.SetFlags(0)
	fi, err = os.Stdin.Stat()
	reject(err)

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		data, err = ioutil.ReadAll(os.Stdin)
		reject(err)
	}

	err = xgoinstall.Run(os.Args[1:], data, os.Stdout, os.Stderr)
	reject(err)
}

func reject(err error) {
	if err != nil && err != flag.ErrHelp {
		log.Println(err)
		exitCode := 1
		if ecoder, ok := err.(interface{ ExitCode() int }); ok {
			exitCode = ecoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}
