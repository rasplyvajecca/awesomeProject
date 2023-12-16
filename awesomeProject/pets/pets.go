package pets

type Pets struct {
	Owner string
	Name  string
}

func AddPetsInfo() []Pets {
	return []Pets{
		{Owner: "Granny", Name: "Cat"},
		{Owner: "Dad", Name: "Dog"},
		{Owner: "Son", Name: "Dog"},
		{Owner: "Mom", Name: "Ret"},
		{Owner: "Mom", Name: "Rabbit"},
	}
}
