package action

import (
	"fmt"
	"io"
	"os"

	"github.com/muesli/crunchy"
	"github.com/urfave/cli"
)

// Check validates password against cracklib
func (s *Action) Check(c *cli.Context) error {
	t, err := s.Store.Tree()
	if err != nil {
		return err
	}

	var out io.Writer
	out = os.Stdout

	foundWeakPasswords := false
	for _, secret := range t.List(0) {
		content, err := s.Store.Get(secret)
		if err != nil {
			return err
		}

		if err = crunchy.ValidatePassword(string(content)); err != nil {
			foundWeakPasswords = true
			fmt.Fprintf(out, "Detected weak password for %s: %v\n", secret, err)
		}
	}

	if !foundWeakPasswords {
		fmt.Fprintln(out, "No weak passwords detected.")
	}

	return nil
}
