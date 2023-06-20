package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

type Monitor struct {
	Title       string
	Description string
	Path        string
}

func ListMonitors(i string) ([]Monitor, error) {
	var monitors []Monitor
	iPath := fmt.Sprintf("./%s/monitors", i)

	files, err := os.ReadDir(iPath)
	if err != nil {
		return monitors, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			m := Monitor{Path: filepath.Join(iPath, file.Name())}
			err := m.processJSONFile()
			if err != nil {
				fmt.Printf("Error processing file %s: %v", m.Path, err)
				continue
			}
			monitors = append(monitors, m)
		}
	}

	return monitors, nil
}

func (m *Monitor) processJSONFile() error {
	var err error

	file, err := os.Open(m.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	var data datadogV1.Monitor

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}

	m.Title = *data.Name
	m.Description = data.Query

	return nil
}
