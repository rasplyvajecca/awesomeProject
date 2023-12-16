package appliance

type Appliance struct {
	Name  string
	Brand string
	Count int
}

func AddApplianceInfo() []Appliance {
	return []Appliance{
		{Name: "Mac", Brand: "Apple", Count: 1},
		{Name: "Iphone X", Brand: "Apple", Count: 1},
		{Name: "Iphone XR", Brand: "Apple", Count: 1},
		{Name: "Iphone 10", Brand: "Apple", Count: 1},
		{Name: "Iphone 10 pro", Brand: "Apple", Count: 1},
	}
}
