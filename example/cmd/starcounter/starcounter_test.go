package main

import (
	"strings"
	"testing"

	"github.com/gregoryv/cmdline"
	"github.com/gregoryv/cmdline/clitest"
)

func TestStarCounter(t *testing.T) {
	okCases := map[string]string{
		"no args":    "",
		"short help": "-h",
		"long help":  "--help",
		"small size": "-size small",
	}
	for name, args := range okCases {
		t.Run(name, func(t *testing.T) {
			checkMain(t, args, 0)
		})
	}

	badCases := map[string]string{
		"heavy weight":     "-weight heavy",
		"unknown argument": "-nosuch",
	}
	for name, args := range badCases {
		t.Run(name, func(t *testing.T) {
			checkMain(t, args, 1)
		})
	}
}

func checkMain(t *testing.T, in string, expectedExitCode int) {
	t.Helper()
	args := strings.Split("test "+in, " ")
	sh := clitest.NewShellT(args...)
	cmdline.DefaultShell = sh
	defer sh.Cleanup()

	main()

	if sh.ExitCode != expectedExitCode {
		t.Error(sh.Dump())
	}
}
