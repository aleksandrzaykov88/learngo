package main

import "fmt"

// getTransport is a factory-method
func getTransport(transportType string) (iTransport, error) {
	if transportType == "seaTransport" {
		return newSeaTransport(), nil
	}
	if transportType == "landTransport" {
		return newLandTransport(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
