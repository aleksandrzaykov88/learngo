package main

import "fmt"

type TapePlayer struct {
	Batteries string
}

type TapeRecorder struct {
	Microphones int
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func playList(device TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
