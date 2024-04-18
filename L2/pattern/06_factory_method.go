package pattern

type action string

const (
	LightSystem action = "LightSystem"
	SoundSysteam action = "SoundSystem"
	ClimateControl action = "ClimateControl"
)

//
type Creator interface {
	CreateProduct(action action) Product 
}



type Product interface {
	Use() string
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}


func (c *ConcreteCreator) CreateProduct(action action) Product {
	var product Product
	switch action {
	case LightSystem:
		product = &LightSystemProduct{string(action)}
	case SoundSysteam:
		product = &SoundSystemProduct{string(action)}
	case ClimateControl:
		product = &ClimateControlProduct{string(action)}
	}
	return product
}

type LightSystemProduct struct {
	action string
}

func (p *LightSystemProduct) Use() string {
	return "Activating: " + p.action
}

type SoundSystemProduct struct {
	action string
}

func (p *SoundSystemProduct) Use() string {
	return "Activating: " + p.action
}

type ClimateControlProduct struct {
	action string
}

func (p *ClimateControlProduct) Use() string {
	return "Activating: " + p.action
}
