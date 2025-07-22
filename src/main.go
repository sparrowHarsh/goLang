package main

import (
	"GoLang/composition"
	"fmt"
)



func main() {
	/* 
	This part will give error due to incapsulation 
	p := pkg.Person{
		firstName: "Vijay",
		lastName:  "Kumar",
		age:       30,
	}

	*/
	
	/*
	p := pkg.Person{}
	p.SetFirstName("Harsh")
	w := p.Walk()
	fmt.Println(w)
	*/

	/* Polymorphism example starts */
	
	// in The left hand side we can have same interface but at the right hand side we can have any struct which implement that class
	/*
	var c polymorphism.Shape = polymorphism.Circle{}
	var s polymorphism.Shape = polymorphism.Square{}

	c.Render()
	s.Render()

	Polymorphism end */

	c := composition.NewCar("a", 600, 300)
	fmt.Print(c);
}
