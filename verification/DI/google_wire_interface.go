// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package DI

import (
	"fmt"

	"github.com/google/wire"
)

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func provideBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

var Set = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideBar)

func InitBar() string {
	wire.Build(Set)
	return ""
}

func Google_Wire_Interface_Exam() {
	fmt.Println("Google_Wire_Interface_Exam: before invoke")
	b := InitBar()
	fmt.Println(b)
}