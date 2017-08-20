package fsutil

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func TestCleanPath(t *testing.T) {
	m := map[string]string{
		filepath.Join("home", "user", "..", "bob", ".password-store"): filepath.Join("home", "bob", ".password-store"),
		filepath.Join("home", "user", "/", ".password-store"):         filepath.Join("home", "user", ".password-store"),
	}
	home, _ := homedir.Dir()
	expanded, _ := homedir.Expand(filepath.Join("~", ".password-store"))
	m[expanded] = filepath.Join(home, ".password-store")

	for in, out := range m {
		got := CleanPath(in)
		if out != got {
			t.Errorf("Mismatch for %s: %s != %s", in, got, out)
		}
	}
}

func TestIsDir(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Failed to create tempdir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tempdir)
	}()
	fn := filepath.Join(tempdir, "foo")
	if err := ioutil.WriteFile(fn, []byte("bar"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	if !IsDir(tempdir) {
		t.Errorf("Should be a dir: %s", tempdir)
	}
	if IsDir(fn) {
		t.Errorf("Should be not dir: %s", fn)
	}
}

func TestIsFile(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Failed to create tempdir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tempdir)
	}()
	fn := filepath.Join(tempdir, "foo")
	if err := ioutil.WriteFile(fn, []byte("bar"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %s", err)
	}
	if IsFile(tempdir) {
		t.Errorf("Should be a dir: %s", tempdir)
	}
	if !IsFile(fn) {
		t.Errorf("Should be not dir: %s", fn)
	}
}

func TestTempdir(t *testing.T) {
	tempdir, err := ioutil.TempDir(tempdirBase(), "gopass-")
	if err != nil {
		t.Fatalf("Failed to create tempdir: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tempdir)
	}()
}

func TestShred(t *testing.T) {
	tempdir, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Failed to create tempdir: %s", err)
	}
	fn := filepath.Join(tempdir, "file")
	fh, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to open file: %s", err)
	}
	buf := make([]byte, 1024)
	for i := 0; i < 10*1024; i++ {
		_, _ = rand.Read(buf)
		_, _ = fh.Write(buf)
	}
	_ = fh.Close()
	if err := Shred(fn, 8); err != nil {
		t.Fatalf("Failed to shred the file: %s", err)
	}
	if IsFile(fn) {
		t.Errorf("Failed still exists after shreding: %s", fn)
	}
	defer func() {
		_ = os.RemoveAll(tempdir)
	}()
}
