package pkg

type Person struct {
	firstName string
	lastName  string
	age       int
}

// Receiver Function
func (p Person) Walk() string {
	return p.firstName + p.lastName + " likes walking"
}

// pointer based receiver function, basically used for setting or changing the value of variables of a struct

// First letter should be capital as it's needed to be public
func (p *Person) SetFirstName(f string) {
	p.firstName = f
}

// Getter Function
func (p *Person) FirstName() string {
	return p.firstName
}


// creating funciton object in different ways, This is kind of constructor
func NewPerson(fn string, ln string, a int)  Person{
	return Person{
		firstName: fn,
		lastName: ln,
		age: a,
	}
}