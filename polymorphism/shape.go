package polymorphism

import "fmt"

/* Interface will contain only functions, This functions will be implemented by struct*/
type Shape interface{
	Render()
}

type Circle struct{}


// This is how interface function is implemented in go
func (C Circle) Render(){
	fmt.Println("Circle is getting rendered")
}

type Square struct{}

func (C Square) Render(){
	fmt.Println("Square is getting rendered")
}
