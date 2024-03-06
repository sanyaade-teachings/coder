//go:build !windows
// +build !windows

// Windows tests fail because the \n\r vs \n. It's not worth trying
// to replace newlines for os tests. If people start using this tool on windows
// and are seeing problems, then we can add build tags and figure it out.
package main

import (
	"testing"

	"github.com/coder/coder/v2/scripts/gentest"
)

func TestTSTypings(t *testing.T) {
	t.Parallel()

	gentest.TestGeneration(t, gentest.TestGenerationParams{
		Generate: func(dir string) (string, error) {
			return Generate(dir)
		},
		TestDataDir:     "testdata",
		GoldenExtension: "ts",
	})
}
