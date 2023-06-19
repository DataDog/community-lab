package main

import (
	"datadog_comunity/src"
)

func main() {

	var integrations = src.CreateDataFile()
	integrations.CreateIntegrationPage()
}
