package sharepoint

import (
	"os"
	"testing"

	"github.com/365admin/cava/config"
)

func TestMain(m *testing.M) {
	config.Setup("../.env")
	code := m.Run()

	os.Exit(code)
}
