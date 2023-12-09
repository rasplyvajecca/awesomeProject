package kids

type Kids struct {
	Member string
	Name   string
	Age    int
}

func AddKidsInfo() []Kids {
	return []Kids{
		{Member: "Друг Даши", Name: "Дима", Age: 21},
		{Member: "Друг Макса", Name: "Игорь", Age: 15},
		{Member: "Друг Макса", Name: "Максим", Age: 16},
		{Member: "Друг Макса", Name: "Дима", Age: 21},
		{Member: "Друг Гриши", Name: "Тима", Age: 2},
	}
}
