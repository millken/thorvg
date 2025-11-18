package thorvg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dnsoa/go/assert"
)

const TEST_DIR = "./resources"

func TestLoadFont(t *testing.T) {
	a := assert.New(t)

	// Test loading valid font
	err := LoadFont(TEST_DIR + "/Arial.ttf")
	a.NoError(err)

	// Test loading invalid font
	err = LoadFont(TEST_DIR + "/invalid.ttf")
	a.Error(err)

	// Test loading with empty path
	err = LoadFont("")
	a.Error(err)

	// Test unloading font
	err = UnloadFont(TEST_DIR + "/Arial.ttf")
	a.NoError(err)

	// Test unloading invalid font
	err = UnloadFont(TEST_DIR + "/invalid.ttf")
	a.Error(err)

	// Test unloading with empty path
	err = UnloadFont("")
	a.Error(err)
}

func TestLoadFontData(t *testing.T) {
	a := assert.New(t)

	// Read test font data
	data, err := os.ReadFile(filepath.Join(TEST_DIR, "Arial.ttf"))
	a.NoError(err)

	// Test loading font data with copy=true
	err = LoadFontData("ArialTest", data, "ttf", true)
	a.NoError(err)

	// Test loading font data with copy=false
	err = LoadFontData("ArialTest2", data, "", false)
	a.NoError(err)

	// Test loading with empty name
	err = LoadFontData("", data, "ttf", false)
	a.Error(err)

	// Test loading with empty data
	err = LoadFontData("Test", []byte{}, "ttf", false)
	a.Error(err)

	// Test loading with unknown mimetype
	err = LoadFontData("ArialUnknown", data, "unknown", false)
	a.NoError(err) // Should succeed as per C API behavior

	// Test unloading loaded fonts
	err = UnloadFont("ArialTest")
	a.Error(err) // Should fail as it's loaded by name, not path

	// Note: UnloadFont is for path-based fonts, not name-based
}
