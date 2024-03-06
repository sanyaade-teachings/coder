package gentest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestGenerationParams struct {
	Generate        func(dir string) (string, error)
	TestDataDir     string
	GoldenExtension string
}

// TestGeneration will perform a generate command and compare it to a golden file
// with the expected output.
func TestGeneration(t *testing.T, params TestGenerationParams) {
	if !strings.HasPrefix(".", params.GoldenExtension) {
		params.GoldenExtension = "." + params.GoldenExtension
	}
	files, err := os.ReadDir(params.TestDataDir)
	require.NoError(t, err, "read dir")

	for _, f := range files {
		if !f.IsDir() {
			// Only test directories
			continue
		}
		f := f
		t.Run(f.Name(), func(t *testing.T) {
			t.Parallel()

			dir := filepath.Join(".", "testdata", f.Name())
			output, err := params.Generate("./" + dir)
			require.NoErrorf(t, err, "generate %q", dir)

			var goldenFile string
			entries, err := os.ReadDir(dir)
			require.NoError(t, err, "read dir")
			for _, entry := range entries {
				if filepath.Ext(entry.Name()) == params.GoldenExtension {
					goldenFile = entry.Name()
					break
				}
			}
			require.NotEmpty(t, goldenFile, "golden file not found")

			golden := filepath.Join(dir, goldenFile)
			expected, err := os.ReadFile(golden)
			require.NoErrorf(t, err, "read file %s", golden)
			expectedString := strings.TrimSpace(string(expected))
			output = strings.TrimSpace(output)
			require.Equal(t, expectedString, output, "matched output")
		})
	}
}
