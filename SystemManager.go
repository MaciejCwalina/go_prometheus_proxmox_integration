package main

import (
	"encoding/json"
	"log"
	"os/exec"
)

type SystemManager struct {
	vms     []VirtualMachine
	nodes   []Node
	storage []Storage
}

func (sm *SystemManager) GetInfrastructureData() ([]map[string]interface{}, error) {
	pCmd := exec.Command("/usr/bin/pvesh", "get", "/cluster/resources", "--output-format", "json-pretty")
	pCmd.Environ()
	bytes, err := pCmd.Output()

	if err != nil {
		log.Fatal("Failed to get output of pvesh due to ", err.Error())
		return nil, err
	}

	var infrastructureData []map[string]interface{}
	err = json.Unmarshal(bytes, &infrastructureData)
	if err != nil {
		log.Fatal("Failed to parse json due to ", err.Error())
		return nil, err
	}

	return infrastructureData, nil
}

func (sm *SystemManager) PopulateSlices(infrastructureData []map[string]interface{}) {
	sm.nodes = nil
	sm.vms = nil
	sm.storage = nil
	var vm VirtualMachine
	var node Node
	var storage Storage
	for _, entry := range infrastructureData {
		typeOfEntry := entry["type"]
		switch typeOfEntry {
		case "qemu":
			bytes, err := json.Marshal(entry) // refactor this to just convert to the vm struct instead of doing this...
			if err != nil {
				return
			}

			json.Unmarshal(bytes, &vm)
			sm.vms = append(sm.vms, vm)
			log.Println("Found a QEMU entry")
		case "node":
			bytes, err := json.Marshal(entry)
			if err != nil {
				return
			}

			json.Unmarshal(bytes, &node)
			sm.nodes = append(sm.nodes, node)
			log.Println("Found a Node Entry")
		case "storage":
			bytes, err := json.Marshal(entry)
			if err != nil {
				return
			}

			json.Unmarshal(bytes, &storage)
			sm.storage = append(sm.storage, storage)
			log.Println("Found a Storage Entry")
		default:
			log.Println("Unkown entry ", typeOfEntry)
		}
	}
}
