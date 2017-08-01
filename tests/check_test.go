package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	ts := newTester(t)
	defer ts.teardown()

	ts.initStore()
	ts.initSecrets("")

	list := `Detected weak password for fixed/secret: Password is too short`
	out, err := ts.run("check")
	assert.NoError(t, err)
	assert.Equal(t, strings.TrimSpace(list), out)
}
