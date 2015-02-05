package testdata

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"v.io/core/veyron/lib/modules"
	"v.io/core/veyron/lib/testutil/integration"
	_ "v.io/core/veyron/profiles"
)

func TestHelperProcess(t *testing.T) {
	modules.DispatchInTest()
}

// redirect redirects the stdout of the given invocation to the file at the
// given path.
func redirect(t *testing.T, inv integration.Invocation, path string) {
	if err := ioutil.WriteFile(path, []byte(inv.Output()), 0600); err != nil {
		t.Fatalf("WriteFile(%q) failed: %v\n", path, err)
	}
}

func TestBlessSelf(t *testing.T) {
	env := integration.New(t)
	defer env.Cleanup()

	var (
		outputDir         = env.TempDir()
		aliceDir          = filepath.Join(outputDir, "alice")
		aliceBlessingFile = filepath.Join(outputDir, "aliceself")
	)

	bin := env.BuildGoPkg("v.io/core/veyron/tools/principal")
	bin.Start("create", aliceDir, "alice").WaitOrDie(os.Stdout, os.Stderr)

	bin = bin.WithEnv([]string{"VEYRON_CREDENTIALS=" + aliceDir})
	redirect(t, bin.Start("blessself", "alicereborn"), aliceBlessingFile)
	got := bin.Start("dumpblessings", aliceBlessingFile).Output()
	want := `Blessings          : alicereborn
PublicKey          : ([0-9a-f]{2}:){15}[0-9a-f]{2}
Certificate chains : 1
Chain #0 \(1 certificates\). Root certificate public key: ([0-9a-f]{2}:){15}[0-9a-f]{2}
  Certificate #0: alicereborn with 0 caveats`
	if regexp.MustCompile(want).FindString(got) == "" {
		t.Fatalf("wanted something matching \n%s, got \n%s\n", want, got)
	}
}
