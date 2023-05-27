package main

import (
	"fmt"

	"github.com/tullur/patterns/behavioral"
	"github.com/tullur/patterns/creational"
	"github.com/tullur/patterns/structural"
)

func main() {
	var (
		id        = "6fa49e0a"
		email     = "test@example.com"
		firstName = "Test"
		lastName  = "Dummy"
	)

	creational.NewCustomer(id, email)
	creational.NewCustomer(id, email, creational.WithName(&firstName, &lastName))
	creational.NewCustomer(
		id,
		email,
		creational.WithName(&firstName, &lastName),
		creational.WithLoyaltyPoints(200),
	)

	singletonInstance := creational.GetSingletoneInstance()
	fmt.Println(singletonInstance)

	requestBuilder := creational.RequestBuilder{}
	requestBuilder.SetHTTPMethod("GET")
	requestBuilder.SetHeaders([]string{"application/json", "application/xml"})
	requestBuilder.SetAuthorization()
	requestBuilder.SetQueryParams([]string{"limit=50"})
	requestBuilder.SetURL("http://localhost:3000/v1/posts")

	request := requestBuilder.Build()
	fmt.Println(request)

	audioPlayer := &structural.AudioPlayer{MediaAdapter: &structural.MediaAdapter{}}
	audioPlayer.Play("mp3", "song.mp3")
	audioPlayer.Play("vlc", "movie.vlc")
	audioPlayer.Play("mp4", "video.mp4")
	audioPlayer.Play("avi", "video.avi")

	file := behavioral.NewParser(&behavioral.CsvFile{})
	file.ParseFile("~/Documents/test.csv")

	airCon := &behavioral.AirCon{}
	remoteControl := &behavioral.RemoteControl{}

	airConOnCommand := &behavioral.OnAirCon{AC: airCon}
	airConOffCommand := &behavioral.OnAirCon{AC: airCon}

	remoteControl.SetOnCommand(airConOnCommand)
	remoteControl.SetOffCommand(airConOffCommand)

	remoteControl.On()
	remoteControl.Off()
}
