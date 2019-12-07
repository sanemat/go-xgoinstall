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
	if err != nil && err != flag.ErrHelp {
		exitWithCode(err)
	}

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil && err != flag.ErrHelp {
			exitWithCode(err)
		}
	}

	err = xgoinstall.Run(os.Args[1:], data, os.Stdout, os.Stderr)
	if err != nil && err != flag.ErrHelp {
		exitWithCode(err)
	}
}

func exitWithCode(err error) {
	log.Println(err)
	exitCode := 1
	if ecoder, ok := err.(interface{ ExitCode() int }); ok {
		exitCode = ecoder.ExitCode()
	}
	os.Exit(exitCode)
}
