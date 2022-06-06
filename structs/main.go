package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

/* Printing alex
func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // print out all the values and the field names
}

*/

/* Printing alex using var
func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // print out all the values and the field names
}

*/

/* Printing jim after updateName without pointer
func main() {
	jim := person{firstName: "Jim", lastName: "Halpert", contactInfo: contactInfo{email: "jim@office.com", zipcode: 18540}}
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName // Doesn't mutate
}
// */

/* Printing jim after updateName using pointer
func main() {
	jim := person{firstName: "Jim", lastName: "Halpert", contactInfo: contactInfo{email: "jim@office.com", zipcode: 18540}}
	jimPointer := &jim // & turn jim to memory address or a pointer and assign to jimPointer
	jimPointer.updateName("Jimmy")
	jim.print()

	slice := []string{"Hi", "There", "How"}
	updateSlice(slice)
	fmt.Println(slice)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

*/

func main() {
	jim := person{firstName: "Jim", lastName: "Halpert", contactInfo: contactInfo{email: "jim@office.com", zipcode: 18540}}
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
