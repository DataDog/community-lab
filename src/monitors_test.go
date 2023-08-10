package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/stretchr/testify/assert"
)

func (m *Monitor) validateSchema(t *testing.T) {
	file, err := os.Open(m.Path)
	assert.Equal(t, err == nil, true, m.Path)
	defer file.Close()

	var data datadogV1.Monitor

	err = json.NewDecoder(file).Decode(&data)
	assert.Equal(t, err == nil, true, "monitor: %s, error: %s", m.Path, err)
}

func TestMonitors(t *testing.T) {
	for _, i := range integrations {
		iPath := fmt.Sprintf("../%s/monitors", i)

		files, err := os.ReadDir(iPath)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
				m := Monitor{Path: filepath.Join(iPath, file.Name())}
				m.validateSchema(t)
			}
		}
	}
}
