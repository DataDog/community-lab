package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

type Dash struct {
	Title       string
	Description string
	Image       string
	Path        string
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
			dash := Dash{Path: filepath.Join(iPath, file.Name())}
			err := dash.processJSONFile()
			if err != nil {
				fmt.Printf("Error processing file %s: %v", dash.Path, err)
				continue
			}
			dash.findImage()
			dashboards = append(dashboards, dash)
		}
	}

	return dashboards, nil
}

func (d *Dash) processJSONFile() error {
	var err error

	file, err := os.Open(d.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	var data datadogV1.Dashboard

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}

	d.Title = data.Title
	d.Description = data.GetDescription()

	return nil
}

func (d *Dash) findImage() {
	imgPath := strings.TrimSuffix(d.Path, "json") + "png"
	if _, err := os.Stat(imgPath); err == nil {
		d.Image = imgPath
	}
}
