// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package DI

import (
	"fmt"

	"github.com/google/wire"
)

type Message string
type Greeter struct {
	Message Message // <- adding a Message field
}
type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

type Msg string
type Eve struct {M Msg}

func NewMessage() Message {
	return Message("Hi there!")
}
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}


func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewMsg() Msg {
	fmt.Println("NewMsg()")
	return Msg("a msg to the world")
}

func NewEve(m Msg) Eve {
	return Eve{M: m}
}

func (e Eve) PrintEve() {
	fmt.Println("carry message: ", e.M)
}

func globalSet() wire.ProviderSet {
	return wire.NewSet(
		NewMsg,
		NewEve,
		NewEvent,
		NewGreeter,
		NewMessage,
	)
}

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}

func InitEve() Eve{
	wire.Build(NewMsg, NewEve)
	return Eve{}
}

func Google_Wire_Struct_Exam() {
	fmt.Println("Google_Wire_Struct_Exam: before invoke")
	e := InitializeEvent()
	e.Start()
}

func Google_Wire_Struct_Global() {
	fmt.Println("Google_Wire_Struct_Global: before invoke")
	e := InitEve()
	e.PrintEve()
}
