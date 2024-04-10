package pattern

import "fmt"

// Интерфейс стратегии
type Strategy interface {
	DoOperation()
}

// Конкретные стратегии
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) DoOperation() {
	fmt.Println("Выполняем сложение")
}

type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) DoOperation() {
	fmt.Println("Выполняем вычитание")
}

// Контекст
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecuteStrategy() {
	if c.strategy != nil {
		c.strategy.DoOperation()
	}
}

func test_07() {
	context := Context{}

	addStrategy := &ConcreteStrategyAdd{}
	context.SetStrategy(addStrategy)
	context.ExecuteStrategy()

	subtractStrategy := &ConcreteStrategySubtract{}
	context.SetStrategy(subtractStrategy)
	context.ExecuteStrategy()
}
