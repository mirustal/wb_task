package pattern

type Visitor interface {
	VisitCarWash(cw *CarWash) string
	VisitRepairShop(rs *RepairShop) string
	VisitRefuelingStation(rs *RefuelingStation) string
}

type Place interface {
	Accept(v Visitor) string
}

type Mechanic struct {
}

func (m *Mechanic) VisitCarWash(cw *CarWash) string {
	return cw.WashCar()
}

func (m *Mechanic) VisitRepairShop(rs *RepairShop) string {
	return rs.RepairCar()
}

func (m *Mechanic) VisitRefuelingStation(rs *RefuelingStation) string {
	return rs.RefuelCar()
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var result string
	for _, place := range c.places {
		result += place.Accept(v) + "\n"
	}
	return result
}

type CarWash struct {
}

func (cw *CarWash) Accept(v Visitor) string {
	return v.VisitCarWash(cw)
}

func (cw *CarWash) WashCar() string {
	return "Мойка автомобиля..."
}

type RepairShop struct {
}


func (rs *RepairShop) Accept(v Visitor) string {
	return v.VisitRepairShop(rs)
}


func (rs *RepairShop) RepairCar() string {
	return "Ремонт автомобиля..."
}


type RefuelingStation struct {
}


func (rs *RefuelingStation) Accept(v Visitor) string {
	return v.VisitRefuelingStation(rs)
}


func (rs *RefuelingStation) RefuelCar() string {
	return "Заправка автомобиля..."
}
