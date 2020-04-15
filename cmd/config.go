package main

import (
	"fmt"
	"strconv"
	"strings"
)

// commands contains list of available commands that can be run from
// cmd package
// current implementation can be extended
var commands = map[string]func(args []string) error {
	"build-word-chain": cmdBuildWordChain,
	//
}

// Errors and outputs
var (
	ErrWrongArgsNum = func(num int) string {
		return fmt.Sprintf("Wrong number of arguments used: %s\n\n", strconv.Itoa(num))
	}
	ErrUnknownCmd = func(cmd string) string {
		return fmt.Sprintf("Unknown command %q\n", cmd)
	}
	ErrBoundsNotSameLength = func() error {
		return fmt.Errorf("Start and end word bounds need to be of the same length")
	}
	ErrEqualValBounds = func(stBound, endBound string) error {
		return fmt.Errorf("Start and end word bounds are the same: %s, %s. No result can be obtained", stBound, endBound)
	}
	ErrBuildShortestPath = func(err string) error {
		return fmt.Errorf("Error creating shotest path: %s", err)
	}
	
	CmdBuildUsageOutput = func (cmd string) string {
		return fmt.Sprintf("Usage of %s: %s -start <starting-word> -end <ending-word>\n", cmd, cmd)
	}
	RunHelpCmdOutput = func (cmd string) string {
		return fmt.Sprintf("Run '%s <command> -help' to learn more about each command.\n", cmd)
	}
	AvailCmdsOutput = func () string {
		return fmt.Sprintf("Available commands are:\n\t%s\n\n", strings.Join(availableCmds(), "\n"))
	}
)
