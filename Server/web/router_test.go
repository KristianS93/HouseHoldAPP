package web

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// initalize server
	// starting db

	os.Exit(m.Run())

	// clear database
}
