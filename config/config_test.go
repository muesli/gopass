package config

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestPwStoreDir(t *testing.T) {
	home, _ := homedir.Dir()
	for in, out := range map[string]string{
		"":        filepath.Join(home, ".password-store"),
		"work":    filepath.Join(home, ".password-store-work"),
		"foo/bar": filepath.Join(home, ".password-store-foo-bar"),
	} {
		got := PwStoreDir(in)
		if got != out {
			t.Errorf("Mismatch for %s: %s != %s", in, got, out)
		}
	}
}
