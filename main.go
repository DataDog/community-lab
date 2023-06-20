package main

import (
	"datadog_community/src"
)

func main() {

	var integrations = src.CreateDataFile()
	integrations.CreateIntegrationPage()
}
