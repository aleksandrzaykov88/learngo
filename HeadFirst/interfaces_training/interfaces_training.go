package main

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

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

//PlayList get music player device and plays songs.
func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

//tryOut tests music player device.
func tryOut(player Player) {
	player.Play("Test Track!")
	player.Stop()
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	p := TapePlayer{}
	r := TapeRecorder{}
	tryOut(p)
	tryOut(r)
}
