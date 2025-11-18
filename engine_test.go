package thorvg_test

import (
	"testing"

	"github.com/dnsoa/go/assert"
	"github.com/millken/thorvg"
)

func TestInitTerm(t *testing.T) {
	r := assert.New(t)
	t.Run("Engine Init Term", func(t *testing.T) {
		r.NoError(thorvg.Init(0))
		r.NoError(thorvg.Term())
	})
}

func TestVersion(t *testing.T) {
	r := assert.New(t)
	major, minor, micro, version, res := thorvg.Version()
	t.Logf("Version: %d.%d.%d, %s", major, minor, micro, version)
	r.True(res.IsSuccess())
}
