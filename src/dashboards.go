package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

type Dash struct {
	Title       string
	Description string
}

func ListIntegrationDashboard(i string) ([]Dash, error) {
	var dashboards []Dash
	iPath := fmt.Sprintf("./%s/dashboards", i)

	files, err := os.ReadDir(iPath)
	if err != nil {
		return dashboards, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(iPath, file.Name())
			dash, err := processJSONFile(filePath)
			if err != nil {
				fmt.Printf("Error processing file %s: %v", filePath, err)
				continue
			}
			dashboards = append(dashboards, dash)
		}
	}

	return dashboards, nil
}

func processJSONFile(filePath string) (Dash, error) {
	var d Dash
	var err error

	file, err := os.Open(filePath)
	if err != nil {
		return d, err
	}
	defer file.Close()

	var data datadogV1.Dashboard

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return d, err
	}

	d.Title = data.Title
	d.Description = data.GetDescription()

	return d, nil
}
