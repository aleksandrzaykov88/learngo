package main

func main() {
	ship, _ := getTransport("seaTransport")
	track, _ := getTransport("landTransport")

	ship.deliver()
	track.deliver()
}
