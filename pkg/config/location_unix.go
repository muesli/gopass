// +build !windows

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/justwatchcom/gopass/pkg/fsutil"
)

// PwStoreDir reads the password store dir from the environment
// or returns the default location ~/.password-store if the env is
// not set
func PwStoreDir(mount string) string {
	if mount != "" {
		return fsutil.CleanPath(filepath.Join(Homedir(), ".password-store-"+strings.Replace(mount, string(filepath.Separator), "-", -1)))
	}
	if d := os.Getenv("PASSWORD_STORE_DIR"); d != "" {
		return fsutil.CleanPath(d)
	}
	return filepath.Join(Homedir(), ".password-store")
}
