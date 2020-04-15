package main

import (
	"fmt"
	"flag"
	"time"
	"github.com/AnNiran/word-chains/build"
)

// cmdBuildWordChain accepts arguments for starting and ending word bounds
// starts bulding word chain if no errros are received
func cmdBuildWordChain(args []string) error {
	// Define command	
	buildWordChainCmd := flag.NewFlagSet("build-word-chain", flag.ExitOnError)

	// Define flags
	startBound := buildWordChainCmd.String("start", "", "Starting word boundary for building path")
	endBound   := buildWordChainCmd.String("end", "", "Ending word boundary for building path")

	// Set up -help output text
	buildWordChainCmd.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Please enter start and end word bounds to build the word chain
Run example:
build-word-chain -start "<start-word>" -end "<end-word>"`,
	)
		buildWordChainCmd.PrintDefaults()
	}

	buildWordChainCmd.Parse(args)

	// We have checked arguments count, no we need to check content

	// Remove all non-alphabetic characters from the input
	trimmed, err := clearNonAlphChars([]string{*startBound, *endBound})
	if err != nil {
		return err
	}

	// Input bounds need to be of the same length and non-equal to each other
	if len(trimmed[0]) != len(trimmed[1]) {
		return ErrBoundsNotSameLength()
	}

	if trimmed[0] == trimmed[1] {
		return ErrEqualValBounds(trimmed[0], trimmed[1])
	}

	// Set up time for measuring execution
	start := time.Now()

	// Run building graph logic
	path, err := build.BuildGraph(trimmed[0], trimmed[1])
	if err != nil {
		// Display execution duration until error
		fmt.Println(measureExecTime(start))
		
		return ErrBuildShortestPath(err.Error())
	}

	// Print the executed time in milliseconds
	fmt.Println(measureExecTime(start))

	// Print the found path
	fmt.Println(path)

	return nil
}
