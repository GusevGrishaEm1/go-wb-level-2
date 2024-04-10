package pattern

import "fmt"

// Интерфейс Создателя
type Creator interface {
	FactoryMethod() Product
}

// Интерфейс объекта
type Object interface {
	Use()
}

// Конкретный объект 1
type ConcreteObject1 struct{}

func (p *ConcreteObject1) Use() {
	fmt.Println("Используем конкретный продукт")
}

// Конкретный объект 2
type ConcreteObject2 struct{}

func (p *ConcreteObject2) Use() {
	fmt.Println("Используем конкретный продукт")
}

// Конкретный создатель 1 для объекта 1
type ConcreteCreator1 struct{}

func (c *ConcreteCreator1) FactoryMethod() Object {
	return &ConcreteObject1{}
}

// Конкретный создатель 2 для объекта 2
type ConcreteCreator2 struct{}

func (c *ConcreteCreator2) FactoryMethod() Object {
	return &ConcreteObject2{}
}

func test_06() {
	// создаем объект 1
	creator1 := &ConcreteCreator1{}
	obj1 := creator1.FactoryMethod()
	obj1.Use()
	// создаем объект 2
	creator2 := &ConcreteCreator2{}
	obj2 := creator2.FactoryMethod()
	obj2.Use()
}
