package xgoinstall

import (
	"flag"
	"fmt"
	"io"
	"log"
)

const cmdName = "x-go-install"

// Run command
func Run(argv []string, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	log.SetPrefix(fmt.Sprintf("[%s] ", cmdName))
	nameAndVer := fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision)
	fs := flag.NewFlagSet(nameAndVer, flag.ContinueOnError)
	fs.SetOutput(errStream)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n", nameAndVer)
		fs.PrintDefaults()
	}
	return nil
}
