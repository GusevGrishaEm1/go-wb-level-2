package pattern

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU Freeze")
}

func (c *CPU) Execute() {
	fmt.Println("CPU Execute")
}

type Memory struct{}

func (m *Memory) Load(addr uint) {
	fmt.Println("Memory Load")
}

type ComputerFacade struct {
	cpu    *CPU
	memory *Memory
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{&CPU{}, &Memory{}}
}

func (cf *ComputerFacade) Start() {
	cf.cpu.Freeze()
	cf.memory.Load(0)
	cf.cpu.Execute()
}

func test_01() {
	computer := NewComputerFacade()
	computer.Start()
}
