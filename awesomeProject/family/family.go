package family

type Family struct {
	Member string
	Name   string
	Age    int
}

func AddFamilyInfo() []Family {
	return []Family{
		{Member: "Mother", Name: "Mom", Age: 20},
		{Member: "Father", Name: "Dad", Age: 25},
		{Member: "Grandmother", Name: "Granny", Age: 50},
		{Member: "Grandfather", Name: "Grandpa", Age: 60},
		{Member: "Uncle", Name: "Uncle", Age: 32},
	}
}
