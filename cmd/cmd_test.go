package main

import (
	"fmt"
	"math"
	//"os"

	"testing"
	"github.com/stretchr/testify/assert"
)

var cmdInit   = "./cmd"
var cmdString = "build-word-chain"
var startFl   = "-start"
var endFl     = "-end"

// Test input values of the same length
func TestCmdBuildWordChainSameArgs(t *testing.T) {
	inputArgsSameValues := []string{
		startFl,
		"value-equals-other-boundary",
		endFl,
		"value-equals-other-boundary",
	}

	expected := ErrEqualValBounds(inputArgsSameValues[1], inputArgsSameValues[3])
	res := cmdBuildWordChain(inputArgsSameValues)

	assert.Equal(t, expected, res)
}

// Test input arguments are not of the same length
func TestCmdBuildWordChainArgsNotSameLength(t *testing.T) {
	err := ErrBoundsNotSameLength()

	testArgs := [][]string{
		{
			"Test Input 1",
			startFl,
			"test-start",
			endFl,
			"test-end",
		},
		{
			"Test input 2",
			startFl,
			"a",
			endFl,
			"abc",
		},
		{
			"Test Input 3",
			startFl,
			"random-test-start-string",
			endFl,
			"random-test-end-string",
		},
		{
			"Test Input 4",
			startFl,
			fmt.Sprintf("a!@#$^&*()_$&346%d", math.MaxInt64),
			endFl,
			"some-very-long-string-input58934^&*%$75389457",
		},
	}

	for _, test := range testArgs {
		res := cmdBuildWordChain(test[1:])
		assert.Equal(t, err, res, fmt.Sprintf("Failed at test %s", test[0])) 
	}
}

/*
// Help function for capturing multiple outputs in the code
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
*/