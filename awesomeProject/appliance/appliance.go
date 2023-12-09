package appliance

type Appliance struct {
	Name  string
	Brand string
	Count int
}

func AddApplianceInfo() []Appliance {
	return []Appliance{
		{Name: "Машина", Brand: "Лада", Count: 1},
		{Name: "Сигнализация", Brand: "Signal", Count: 1},
		{Name: "Холодильник", Brand: "Apple", Count: 3},
		{Name: "Микроволновка", Brand: "Bosch", Count: 2},
		{Name: "Робот-пылесос", Brand: "Bosch", Count: 2},
	}
}
