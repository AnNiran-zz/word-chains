package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	"errors"

	"regexp"
	"time"
)

// checkArgsCount checks if proper number of arguments is received from input
// if not - outputs corresponding outputs with information
// about how to properly run the functionality
func checkArgsCount(args []string) error {
	// os.Args need to be 6
	// 1. ./cmd
	// 2. build-word-chain
	// 3. -start
	// 4. <start-word>
	// 5. -end
	// 6. <end-word>
	// but we also have to consider cases where -help flag is called - arguments number is 3
	
	// This condition handles the cases when "-help" flag is called with a command
	// and also is handled in the next condition if the third argument is something else
	if len(args) == 3 && args[2] == "-help" { // TODO: make it better
		return nil
	}

	// Handle all cases when arguments are not 6, we have already checked for "-help"
	if len(args) != 6 {
		fmt.Fprintf(os.Stderr, ErrWrongArgsNum(len(args)-1))
		fmt.Fprintf(os.Stderr, AvailCmdsOutput())
		
		fmt.Fprintf(os.Stderr, RunHelpCmdOutput(args[0]))
	
		// This error content is not used in current implementation
		// Returned value is checked if nil on the other side
		return errors.New(ErrWrongArgsNum(len(args)-1))
	}

	return nil
}

// availableCmds lists the available commands
func availableCmds() []string {
	cmds := make([]string, 0, len(commands))
	for name := range commands {
		cmds = append(cmds, name)
	}

	sort.Strings(cmds)
	return cmds
}

// clearNonAlphChars trims all non-alphabetic characters from words
// except words containing dashes
func clearNonAlphChars(srcBound []string) ([]string, error) {
	// Prepare regular expression for keeping only alphabetic characters
	reg, err := regexp.Compile("[^a-zA-Z-]+")
	if err != nil {
		return nil, err
	}

	for i, _ := range srcBound {
		// Clear all non-alphabetic characters except dashes
		cleared := reg.ReplaceAllString(srcBound[i], "")

		// Trim dashes, if any, from string ends
		trimmed := strings.Trim(cleared, "-")
		srcBound[i] = trimmed
	}

	return srcBound, nil
}

// measureExecTime returns the duration between received time and current moment
func measureExecTime(start time.Time) time.Duration {
	return time.Since(start)
}
