package main

import "fmt"

func main() {
	secretAgent{
		fPerson: person{
			name: "Leo",
			age:  45,
		},
		licenseToKill: true,
	}.speak()

}

func (p secretAgent) speak() {
	fmt.Println(p.fPerson, "you can kill?", p.licenseToKill)
}

type secretAgent struct {
	fPerson       person
	licenseToKill bool
}

type person struct {
	name string
	age  int
}
