package pattern
/*
	Паттерн Builder относится к порождающим паттернам уровня объекта.
	Паттерн Builder определяет процесс поэтапного построения сложного продукта. После того
как будет построена последняя его часть, продукт можно использовать.

*/


import "fmt"


type Director struct {
	builder Builder
}


func (d *Director) Construct() {
	d.builder.MakeEngine("V8 engine")
	d.builder.MakeChassis("SUV chassis")
	d.builder.MakeInterior("Leather interior")
}


type Builder interface {
	MakeEngine(str string)
	MakeChassis(str string)
	MakeInterior(str string)
}


type CarBuilder struct {
	car *Car
}


func (b *CarBuilder) MakeEngine(str string) {
	b.car.Engine = str
}


func (b *CarBuilder) MakeChassis(str string) {
	b.car.Chassis = str
}


func (b *CarBuilder) MakeInterior(str string) {
	b.car.Interior = str
}


type Car struct {
	Engine   string
	Chassis  string
	Interior string
}


func (c *Car) Show() string {
	return fmt.Sprintf("Car with %s, %s, and %s", c.Engine, c.Chassis, c.Interior)
}

