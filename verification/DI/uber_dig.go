package DI

import (
	"fmt"

	"go.uber.org/dig"
)

type A string
type B struct{ArgA A}
type C struct{ArgA A}
type D struct{
	ArgB B
	ArgC C
}
type E struct{}

func NewD(b B, c C) D {
	fmt.Println("NewD()")
	return D{ArgB: b, ArgC: c}
}
func NewB(a A) B {
	fmt.Println("NewB()")
	return B{ArgA: a}
}
func NewC(a A) C {
	fmt.Println("NewC()")
	return C{ArgA: a}
}
func NewA() A {
	fmt.Println("NewA()")
	return A("hello world")
}
func NewE() E{
	fmt.Println("NewE()")
	return E{}
}
func (d D) PrintD() {
	fmt.Printf("print D with B: %v, C: %v\n", d.ArgB, d.ArgC)
}
func (e E) PrintE() {
	fmt.Println("print E")
}

func global() *dig.Container {
	// 创建 dig 对象
	digObj := dig.New()
	// 利用 Provide 注入依赖
	digObj.Provide(NewA)
	digObj.Provide(NewC)
	digObj.Provide(NewB)
	digObj.Provide(NewD)
	digObj.Provide(NewE)

	return digObj
}

func Uber_Dig_Exam() {
	digObj := global()

	fmt.Println("before invoke")
	// 根据提前注入的依赖来生成对象
	if err := digObj.Invoke(func(argD D) {
		argD.PrintD()
	}); err != nil {
		panic(err)
	}

	var e E
	if err := digObj.Invoke(func(argE E) {
		e = argE
	}); err != nil {
		panic(err)
	}
	e.PrintE()
}