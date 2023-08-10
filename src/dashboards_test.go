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

func (d *Dash) validateSchema(t *testing.T) {
	file, err := os.Open(d.Path)
	assert.Equal(t, err == nil, true, d.Path)
	defer file.Close()

	var data datadogV1.Dashboard

	err = json.NewDecoder(file).Decode(&data)
	assert.Equal(t, err == nil, true, "dashboars: %s, error: %s", d.Path, err)
}

func TestDashboards(t *testing.T) {
	for _, i := range integrations {
		iPath := fmt.Sprintf("../%s/dashboards", i)

		files, err := os.ReadDir(iPath)
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
				d := Dash{Path: filepath.Join(iPath, file.Name())}
				d.validateSchema(t)
			}
		}
	}
}
