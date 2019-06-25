package main

import "fmt"

func main() {
	say(
		secretAgent{
			fPerson: person{
				name: "Leo",
				age:  45,
			},
			licenseToKill: true,
		},
	)
}

func say(h human) {
	h.speak()
}

func (p secretAgent) speak() {
	fmt.Println(p.fPerson, "you can kill?", p.licenseToKill)
}

type human interface {
	speak()
}

type secretAgent struct {
	fPerson       person
	licenseToKill bool
}

type person struct {
	name string
	age  int
}
