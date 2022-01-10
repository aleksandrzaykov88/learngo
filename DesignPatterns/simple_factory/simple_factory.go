package main

import "fmt"

type iTransport interface {
	setName(name string)
	getName() string
	deliver()
}

type transport struct {
	name  string
	speed int
}

func (t *transport) setName(name string) {
	t.name = name
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) deliver() {
	fmt.Println("Cargo has been delivered by " + t.getName())
}

type track struct {
	transport
}

func newLandTransport() iTransport {
	return &track{
		transport: transport{
			name:  "truck",
			speed: 4,
		},
	}
}

type ship struct {
	transport
}

func newSeaTransport() iTransport {
	return &ship{
		transport: transport{
			name:  "ship",
			speed: 3,
		},
	}
}

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

func main() {
	ship, _ := getTransport("seaTransport")
	track, _ := getTransport("landTransport")

	ship.deliver()
	track.deliver()
}
