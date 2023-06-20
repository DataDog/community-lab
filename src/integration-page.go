package src

import (
	"fmt"
	"os"
	"text/template"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

var reserved_directories = []string{"src", "_web", ".git", ".github"}

type Integration struct {
	Name       string
	Dashboards []Dash
	Monitors   []Monitor
}

type IntegrationList struct {
	Integrations []Integration
}

func cleanup() {
	os.RemoveAll("./_web/integrations/")
	os.RemoveAll("./_web/_data/")
	os.MkdirAll("./_web/integrations/", 0744)
	os.MkdirAll("./_web/_data/", 0744)
}

func Build() {
	cleanup()
	list := createDataFile()
	list.createIntegrationPages()
}

func createDataFile() IntegrationList {
	var list IntegrationList

	items, _ := os.ReadDir(".")
	for _, item := range items {
		if item.IsDir() && !slices.Contains(reserved_directories, item.Name()) {
			integra := Integration{Name: item.Name()}

			dashs, err := ListIntegrationDashboard(item.Name())
			if err != nil {
				fmt.Println(err)
			} else {
				integra.Dashboards = dashs
			}
			monitors, err := ListMonitors(item.Name())
			if err != nil {
				fmt.Println(err)
			} else {
				integra.Monitors = monitors
			}

			yamlData, err := yaml.Marshal(&integra)
			if err != nil {
				fmt.Printf("Error while Marshaling. %v", err)
				continue
			}
			dataPath := fmt.Sprintf("./_web/_data/integration_%s.yaml", integra.Name)
			err = os.WriteFile(dataPath, yamlData, 0644)
			if err != nil {
				fmt.Printf("Error while Writing. %v", err)
			}
			list.Integrations = append(list.Integrations, integra)

		}
	}
	return list
}

func (il *IntegrationList) createIntegrationPages() {
	t, err := template.ParseFiles("src/integration-template.md")
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, i := range il.Integrations {

		f, err := os.Create(fmt.Sprintf("./_web/integrations/%s.md", i.Name))
		if err != nil {
			fmt.Println("create file: ", err)
			return
		}
		defer f.Close()

		err = t.Execute(f, i)
		if err != nil {
			fmt.Print("execute: ", err)
			return
		}
	}
}
