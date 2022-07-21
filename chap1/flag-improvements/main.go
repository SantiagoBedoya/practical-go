package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes   int
	printUsage bool
	name       string
}

var errInvalidPosArgSpecified = errors.New("more than one positional argument specified")
var usageString = fmt.Sprintf(`Usage: %s <integer> [-h]--help

A greeter app which prints the name you entered <integer> number of times
`, os.Args[0])

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&c.numTimes, "n", 0, "number of times to greet")
	if err := fs.Parse(args); err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errInvalidPosArgSpecified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", c.name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}
	greetUser(c, w)
	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the enter key when done.\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
}

func main() {
	c, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		if errors.Is(err, errInvalidPosArgSpecified) {
			fmt.Fprintln(os.Stdout, err)
		}
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
