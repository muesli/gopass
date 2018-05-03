package config

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	apppaths "github.com/muesli/go-app-paths"
)

// Homedir returns the users home dir or an empty string if the lookup fails
func Homedir() string {
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		return hd
	}

	hd, err := homedir.Dir()
	if err != nil {
		if debug {
			fmt.Printf("[DEBUG] Failed to get homedir: %s\n", err)
		}
		return ""
	}
	return hd
}

// configLocation returns the location of the config file
// (a YAML file that contains values such as the path to the password store)
func configLocation() string {
	// First, check for the "GOPASS_CONFIG" environment variable
	if cf := os.Getenv("GOPASS_CONFIG"); cf != "" {
		return cf
	}

	scope := apppaths.NewScope(apppaths.User, "gopass", "gopass")
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		scope = apppaths.NewCustomHomeScope(hd, "gopass", "gopass")
	}

	cf, err := scope.ConfigPath("config.yml")
	if err != nil {
		return filepath.Join(Homedir(), ".config", "gopass", "config.yml")
	}

	return cf
}

// configLocations returns the possible locations of gopass config files,
// in decreasing priority
func configLocations() []string {
	l := []string{}
	if cf := os.Getenv("GOPASS_CONFIG"); cf != "" {
		l = append(l, cf)
	}

	scope := apppaths.NewScope(apppaths.User, "gopass", "gopass")
	if hd := os.Getenv("GOPASS_HOMEDIR"); hd != "" {
		scope = apppaths.NewCustomHomeScope(hd, "gopass", "gopass")
	}
	if cf, err := scope.ConfigPath("config.yml"); err == nil {
		l = append(l, cf)
	}

	l = append(l, filepath.Join(Homedir(), ".config", "gopass", "config.yml"))
	l = append(l, filepath.Join(Homedir(), ".gopass.yml"))
	return l
}

// Directory returns the configuration directory for the gopass config file
func Directory() string {
	return filepath.Dir(configLocation())
}
