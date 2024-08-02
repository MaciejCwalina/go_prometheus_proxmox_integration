package main

import (
	"log"
)

func main() {
	systemManager := SystemManager{}
	infraData, err := systemManager.GetInfrastructureData()
	if err != nil {
		log.Fatal("Failed to Get Infrastructure Information", err.Error())
		return
	}

	systemManager.PopulateSlices(infraData)
}
