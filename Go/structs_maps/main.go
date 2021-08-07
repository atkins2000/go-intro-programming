package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo //or just contactInfo which saves the var as contactInfo with type contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//yash := person{firstName: "Yash", lastName: "Gulati", contact:contactInfo{email: "divyashg@gmail.com", zipCode: 121001,},}
	//or
	var yash person //default for this declaration is empty string, 0 for float and int and false for bool
	yash.firstName = "Yash"
	yash.lastName = "Gulati"
	yash.contact.email = "divyashg@gmail.com"
	yash.contact.zipCode = 121001
	//yashPointer := &yash //required to make changes here and not in its copy
	//yashPointer.updateName("yashi")
	//or
	yash.updateName("yashi") //automatically converts struct to struct pointer if the function takes only pointer

	//map declaration
	colors := map[string]string{
		"white": "#FFFFFF",
		"black": "#000000",
		"some":  "some",
	}
	//or
	//var colors map[string]string
	//or
	//colors := make(map[string]string)
	//colors["white"] = "#FFFFFF"
	delete(colors, "white")
	printMap(colors)
	//fmt.Println(colors) //prints the map with key value pairs

	yash.print()
}

func (p person) print() {
	fmt.Printf("%+v", p) //prints with fields
}

//func (p person) updateName(newFirstName string) {
//	p.firstName = newFirstName
//} //does not work because go is a pass by value language and the p is a different ram location which holds a copy of the call

func (pointerToName *person) updateName(newFirstName string) {
	(*pointerToName).firstName = newFirstName //*something converts to value
}

//a slice is a DS which houses a ptr to head of an array, the capacity and length(current elements in the array)
//when we make a slice a slice DS and an array is created

//when we update something in a slice with function, we modify the same array because of the ptr to head
//slice, maps, channels are reference type stuff
//int, float, string, bool, structs, arrays are value type

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for", color, "is", hex)
	}
}
