package pattern

import "fmt"

// Command interface
type Command interface {
	Execute()
}

// Light struct
type Light struct {
}

// TurnOnCommand struct
type TurnOnCommand struct {
	light *Light
}

// TurnOffCommand struct
type TurnOffCommand struct {
	light *Light
}

// NewTurnOnCommand returns a new TurnOnCommand instance
func NewTurnOnCommand(light *Light) *TurnOnCommand {
	return &TurnOnCommand{light: light}
}

// NewTurnOffCommand returns a new TurnOffCommand instance
func NewTurnOffCommand(light *Light) *TurnOffCommand {
	return &TurnOffCommand{light: light}
}

func (c *TurnOnCommand) Execute() {
	fmt.Println("Turning on the light")
}

func (c *TurnOffCommand) Execute() {
	fmt.Println("Turning off the light")
}

func test_04() {
	light := &Light{}

	// Создаем команды
	turnOnCommand := NewTurnOnCommand(light)
	turnOffCommand := NewTurnOffCommand(light)

	// Выполняем команды
	turnOnCommand.Execute()
	turnOffCommand.Execute()
}
