package furniture

type Furniture struct {
	Name string
	Size float64
}

func AddFurnitureInfo() []Furniture {
	return []Furniture{
		{Name: "Bed", Size: 2},
		{Name: "Bed", Size: 3},
		{Name: "Chair", Size: 8},
		{Name: "Table", Size: 8},
		{Name: "Chair", Size: 8},
	}
}
