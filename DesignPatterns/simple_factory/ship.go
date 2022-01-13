package main

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
