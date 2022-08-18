package web

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// initalize server
	// starting db

	code := m.Run()
	os.Exit(code)

	// clear database
}
