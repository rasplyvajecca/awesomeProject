package pets

type Pets struct {
	Name  string
	Size  int
	Count int
}

func AddPetsInfo() []Pets {
	return []Pets{
		{Name: "Футболка", Size: 54, Count: 10},
		{Name: "Толстовка", Size: 52, Count: 5},
		{Name: "Джинсы", Size: 34, Count: 2},
		{Name: "Носки", Size: 40, Count: 20},
		{Name: "Куртка", Size: 58, Count: 1},
	}
}
