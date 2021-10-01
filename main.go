package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("FATAL: ", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	printMinified := flag.Bool("m", false, "should print minified")
	flag.Parse()

	bytes, err := getBytesFromArgs(os.Args)
	if err != nil {
		return err
	}

	if !*printMinified {
		var obj map[string]interface{}
		err := json.Unmarshal(bytes, &obj)
		if err != nil {
			return errors.Wrap(err, "unmarshalling")
		}

		prettyBytes, err := json.MarshalIndent(obj, "", "    ")
		if err != nil {
			return errors.Wrap(err, "trying to make pretty")
		}

		err = json.Unmarshal(prettyBytes, &obj)
		if err != nil {
			return errors.Wrap(err, "unmarshalling")
		}

		bytes = prettyBytes
	}

	fmt.Println(string(bytes))

	return nil
}

func getBytesFromArgs(osArgs []string) ([]byte, error) {
	// index 0 is the program name program name
	sBytes := osArgs[1:]

	if len(sBytes) == 0 {
		return nil, errors.New("no bytes passed in")
	}

	// get rid of possible flag
	if sBytes[0] == "-m" {
		sBytes = sBytes[1:]
	}

	// trim possible [] on first and last
	sBytes[0] = strings.Trim(sBytes[0], "[")
	sBytes[len(sBytes)-1] = strings.Trim(sBytes[len(sBytes)-1], "]")

	bytes := make([]byte, len(sBytes))
	for i, sByte := range sBytes {
		iByte, err := strconv.Atoi(sByte)
		if err != nil {
			return nil, errors.Wrapf(err, "turning string '%s' to a int", sByte)
		}

		bytes[i] = byte(iByte)
	}

	return bytes, nil
}
