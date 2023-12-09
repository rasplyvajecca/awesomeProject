package family

type Family struct {
	Member string
	Name   string
	Age    int
}

func AddFamilyInfo() []Family {
	return []Family{
		{Member: "Мама", Name: "Ирина", Age: 41},
		{Member: "Папа", Name: "Даня", Age: 30},
		{Member: "Дочка", Name: "Даша", Age: 20},
		{Member: "Сын", Name: "Максим", Age: 10},
		{Member: "Сын", Name: "Гриша", Age: 2},
	}
}
