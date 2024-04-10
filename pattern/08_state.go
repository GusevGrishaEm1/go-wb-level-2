package pattern

import "fmt"

// Интерфейс состояния
type State interface {
	DoAction(obj *ObjWithState)
}

// Реализации конкретных состояний
type ConcreteStateA struct{}

func (sa *ConcreteStateA) DoAction(obj *ObjWithState) {
	fmt.Println("Состояние A")
	obj.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (sb *ConcreteStateB) DoAction(obj *ObjWithState) {
	fmt.Println("Состояние B")
	obj.SetState(&ConcreteStateA{})
}

// Объект с состоянием
type ObjWithState struct {
	state State
}

func (c *ObjWithState) SetState(s State) {
	c.state = s
}

func (c *ObjWithState) Request() {
	c.state.DoAction(c)
}

func test_08() {
	obj := ObjWithState{state: &ConcreteStateA{}}
	obj.Request()
	obj.SetState(&ConcreteStateB{})
	obj.Request()
}
