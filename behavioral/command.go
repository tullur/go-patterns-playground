package behavioral

import "log"

type Command interface {
	Execute()
}

type AirCon struct{}

func (ac *AirCon) On() {
	log.Println("Air conditioning turned on")
}
func (ac *AirCon) Off() {
	log.Println("Air conditioning turned off")
}

type OnAirCon struct {
	AC *AirCon
}

type OffAirCon struct {
	AC *AirCon
}

func (command *OnAirCon) Execute() {
	command.AC.On()
}

func (command *OffAirCon) Execute() {
	command.AC.Off()
}

type RemoteControl struct {
	onCommand  Command
	offCommand Command
}

func (rc *RemoteControl) SetOnCommand(command Command) {
	rc.onCommand = command
}

func (rc *RemoteControl) SetOffCommand(command Command) {
	rc.offCommand = command
}

func (rc *RemoteControl) On() {
	rc.onCommand.Execute()
}

func (rc *RemoteControl) Off() {
	rc.offCommand.Execute()
}
