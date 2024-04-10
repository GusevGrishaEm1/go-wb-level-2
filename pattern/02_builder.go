package pattern

import "fmt"

// Структура для представления объекта, который мы будем строить
type Product struct {
	Part1 string
	Part2 string
}

// Интерфейс Строителя
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetProduct() Product
}

// Конкретный строитель
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.Part1 = "Part 1 built"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.Part2 = "Part 2 built"
}

func (b *ConcreteBuilder) GetProduct() Product {
	return b.product
}

func test_02() {
	builder := &ConcreteBuilder{}
	builder.BuildPart1()
	builder.BuildPart2()

	product := builder.GetProduct()

	fmt.Println("Product Part 1:", product.Part1)
	fmt.Println("Product Part 2:", product.Part2)
}
