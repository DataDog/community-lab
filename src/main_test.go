package src

import (
	"os"
	"testing"

	"golang.org/x/exp/slices"
)

var integrations []string

func listIntegrations() {
	items, _ := os.ReadDir("..")
	for _, item := range items {
		if item.IsDir() && !slices.Contains(reserved_directories, item.Name()) {
			integrations = append(integrations, item.Name())
		}
	}
}

func TestMain(m *testing.M) {
	listIntegrations()
	code := m.Run()
	os.Exit(code)
}
