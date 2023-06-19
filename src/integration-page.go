package src

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

var reserved_directories = []string{"src", "_web", ".git"}

type Integration struct {
	Name       string
	Dashboards []Dash
}

type IntegrationList struct {
	Integrations []string
}

func CreateDataFile() IntegrationList {
	var list IntegrationList

	items, _ := os.ReadDir(".")
	for _, item := range items {
		if item.IsDir() && !slices.Contains(reserved_directories, item.Name()) {
			list.Integrations = append(list.Integrations, item.Name())

			integra := Integration{Name: item.Name()}

			dashs, err := ListIntegrationDashboard(item.Name())
			if err != nil {
				fmt.Println(err)
			} else {
				integra.Dashboards = dashs
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
		}
	}
	yamlData, err := yaml.Marshal(list)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}
	err = os.WriteFile("./_web/_data/IntegrationList.yaml", yamlData, 0644)
	if err != nil {
		fmt.Printf("Error while Writing. %v", err)
	}
	return list
}

func (il *IntegrationList) CreateIntegrationPage() {
	for _, integ := range il.Integrations {
		err := os.WriteFile(fmt.Sprintf("./_web/integrations/%s.md", integ), []byte(NewIntegrationTemplate(integ)), 0644)
		if err != nil {
			fmt.Printf("Error while Writing. %v", err)
		}
	}
}
