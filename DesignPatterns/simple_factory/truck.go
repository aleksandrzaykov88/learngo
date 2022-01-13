package main

type truck struct {
	transport
}

func newLandTransport() iTransport {
	return &truck{
		transport: transport{
			name:  "truck",
			speed: 4,
		},
	}
}
