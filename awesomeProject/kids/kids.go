package kids

type Kids struct {
	Member string
	Name   string
	Age    int
}

func AddKidsInfo() []Kids {
	return []Kids{
		{Member: "Baby boy", Name: "Ivan", Age: 1},
		{Member: "Baby girl", Name: "Dasha", Age: 5},
		{Member: "Son", Name: "Misha", Age: 11},
		{Member: "Daughter", Name: "Liza", Age: 11},
		{Member: "Cousen", Name: "Maksim", Age: 23},
	}
}
