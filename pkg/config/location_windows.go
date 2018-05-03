// +build windows

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/justwatchcom/gopass/pkg/fsutil"
	apppaths "github.com/muesli/go-app-paths"
)

// PwStoreDir reads the password store dir from the environment
// or returns the default location %localappdata%/gopass/password-store
// if the env is not set
func PwStoreDir(mount string) string {
	if mount != "" {
		return fsutil.CleanPath(filepath.Join(Homedir(), ".password-store-"+strings.Replace(mount, string(filepath.Separator), "-", -1)))
	}
	if d := os.Getenv("PASSWORD_STORE_DIR"); d != "" {
		return fsutil.CleanPath(d)
	}

	scope := apppaths.NewScope(apppaths.User, "gopass", "gopass")
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		scope = apppaths.NewCustomHomeScope(hd, "gopass", "gopass")
	}

	pd, err := scope.DataDir()
	if err != nil {
		return filepath.Join(Homedir(), ".password-store")
	}
	return filepath.Join(pd, "password-store")
}
