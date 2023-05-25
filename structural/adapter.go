package structural

import "fmt"

// Target interface
type MediaPlayer interface {
	Play(fileType string, fileName string)
}

// Adaptee interface
type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

// Adaptee type: VlcPlayer
type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(fileName string) {
	fmt.Println("Playing vlc file:", fileName)
}

func (v *VlcPlayer) PlayMp4(fileName string) {}

// Adaptee type: Mp4Player
type Mp4Player struct{}

func (m *Mp4Player) PlayVlc(fileName string) {}

func (m *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing mp4 file:", fileName)
}

// Adapter type: MediaAdapter
type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (a *MediaAdapter) Play(fileType string, fileName string) {
	if fileType == "vlc" {
		a.advancedMediaPlayer.PlayVlc(fileName)
	} else if fileType == "mp4" {
		a.advancedMediaPlayer.PlayMp4(fileName)
	}
}

// Concrete adapter type: AudioPlayer
type AudioPlayer struct {
	MediaAdapter MediaPlayer
}

func (a *AudioPlayer) Play(fileType string, fileName string) {
	if fileType == "mp3" {
		fmt.Println("Playing mp3 file:", fileName)
	} else if fileType == "vlc" || fileType == "mp4" {
		a.MediaAdapter.Play(fileType, fileName)
	} else {
		fmt.Println("Unsupported file type:", fileType)
	}
}
