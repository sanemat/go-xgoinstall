package main

import (
	"flag"
	"github.com/sanemat/go-xgoinstall"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	err := xgoinstall.Run(os.Args[1:], os.Stdout, os.Stderr)
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
