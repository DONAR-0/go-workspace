package encap

type Employee struct {
	name string
	age  int
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetAge() int {
	return e.age
}

func (e *Employee) SetAge(age int) {
	e.age = age
}
