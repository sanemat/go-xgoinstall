package xgoinstall

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	"golang.org/x/xerrors"
)

const cmdName = "x-go-install"

// Run command
func Run(argv []string, data []byte, outStream, errStream io.Writer) error {
	log.SetOutput(errStream)
	log.SetPrefix(fmt.Sprintf("[%s] ", cmdName))
	nameAndVer := fmt.Sprintf("%s (v%s rev:%s)", cmdName, version, revision)
	fs := flag.NewFlagSet(nameAndVer, flag.ContinueOnError)
	fs.SetOutput(errStream)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n", nameAndVer)
		fs.PrintDefaults()
	}

	var (
		ver             = fs.Bool("version", false, "display version")
		nullTerminators = fs.Bool("0", false, "use NULs as input field terminators")
	)

	if err := fs.Parse(argv); err != nil {
		return err
	}
	if *ver {
		return printVersion(outStream)
	}

	argv = fs.Args()
	if len(argv) >= 1 {
		return xerrors.New("We have no subcommand")
	}

	if len(data) == 0 {
		return nil
	}

	var targets []string
	if *nullTerminators {
		targets = strings.Split(string(data), "\x00")
	} else {
		targets = strings.Fields(string(data))
	}
	for _, v := range targets {
		cmd := exec.Command("go", "install", v)
		stdout, err1 := cmd.Output()
		if err1 != nil {
			return err1
		}
		_, err2 := fmt.Fprint(outStream, string(stdout))
		if err2 != nil {
			return err2
		}
	}

	return nil
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "%s v%s (rev:%s)\n", cmdName, version, revision)
	return err
}
