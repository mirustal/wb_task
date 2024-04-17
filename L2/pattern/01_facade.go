package pattern

/*
Паттерн Facade относится к структурным паттернам уровня объекта.
Паттерн Facade предоставляет высокоуровневый унифицированный интерфейс в виде
набора имен методов к набору взаимосвязанных классов или объектов некоторой
подсистемы, что облегчает ее использование.
Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а
также помогает максимально снизить зависимости одной подсистемы от другой
*/
import "strings"

func NewCar() *CarFacade {
	return &CarFacade{
		engine:    &Engine{},
		chassis:   &Chassis{},
		electronics: &Electronics{},
	}
}

type CarFacade struct {
	engine      *Engine
	chassis     *Chassis
	electronics *Electronics
}


func (c *CarFacade) Operate() string {
	result := []string{
		c.engine.Start(),
		c.chassis.Assemble(),
		c.electronics.Activate(),
	}
	return strings.Join(result, "\n")
}


type Engine struct {
}

// Start запускает двигатель.
func (e *Engine) Start() string {
	return "Engine started"
}


type Chassis struct {
}


func (c *Chassis) Assemble() string {
	return "Chassis assembled"
}


type Electronics struct {
}


func (e *Electronics) Activate() string {
	return "Electronics activated"
}
