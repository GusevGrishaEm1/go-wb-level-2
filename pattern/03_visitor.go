package pattern

import "fmt"

type Transport interface {
	Accept(Visitor)
}

type Car struct {
	Name string
}

func (c *Car) Accept(v Visitor) {
	v.VisitCar(c)
}

type Motorcycle struct {
	Name string
}

func (m *Motorcycle) Accept(v Visitor) {
	v.VisitMotorcycle(m)
}

type Bicycle struct {
	Name string
}

func (b *Bicycle) Accept(v Visitor) {
	v.VisitBicycle(b)
}

type Visitor interface {
	VisitCar(*Car)
	VisitMotorcycle(*Motorcycle)
	VisitBicycle(*Bicycle)
}

type YearCheckVisitor struct{}

func (v *YearCheckVisitor) VisitCar(c *Car) {
	fmt.Printf("Checking car %s for roadworthiness...n", c.Name)
}

func (v *YearCheckVisitor) VisitMotorcycle(m *Motorcycle) {
	fmt.Printf("Checking motorcycle %s for roadworthiness...n", m.Name)
}

func (v *YearCheckVisitor) VisitBicycle(b *Bicycle) {
	fmt.Printf("Checking bicycle %s for roadworthiness...n", b.Name)
}

func test_03() {
	car := &Car{Name: "Toyota"}
	motorcycle := &Motorcycle{Name: "Harley Davidson"}
	bicycle := &Bicycle{Name: "Trek"}

	visitor := &YearCheckVisitor{}

	car.Accept(visitor)
	motorcycle.Accept(visitor)
	bicycle.Accept(visitor)
}
