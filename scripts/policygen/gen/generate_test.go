package gen_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/coder/coder/v2/scripts/gentest"
	"github.com/coder/coder/v2/scripts/policygen/gen"
)

func TestZedPolicyGen(t *testing.T) {
	if runtime.GOOS == "windows" {
		// Windows tests fail because the \n\r vs \n. It's not worth trying
		// to replace newlines for os tests. If people start using this tool on windows
		// and are seeing problems, then we can add build tags and figure it out.
		t.Skip("Skipping on windows")
	}

	t.Parallel()

	const dir = "testdata"
	gentest.TestGeneration(t, gentest.TestGenerationParams{
		Generate: func(dir string) (string, error) {
			src, err := gentest.FindFileWithExtension(dir, ".zed")
			if err != nil {
				return "", err
			}
			schema, err := os.ReadFile(filepath.Join(dir, src))
			if err != nil {
				return "", err
			}
			return gen.Generate(string(schema), gen.GenerateOptions{})
		},
		TestDataDir:     dir,
		GoldenExtension: "go",
	})
}
