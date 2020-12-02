package lesson2

import "fmt"

// Person ...
type Person struct {
	Name string
	Age  int
}

// Account ...
type Account struct {
	ID   int
	Name string
	Person
}

// MySlice test
type MySlice []int

// Add int number to MySlice
func (slc *MySlice) Add(value int) {
	*slc = append(*slc, value)
}

// Count returns number of length
func (slc *MySlice) Count() int {
	return len(*slc)
}

// UnpadeName Unpade name for Person struct
func (p Person) UnpadeName(name string) {
	p.Name = name
}

// SetName Set name for link Person strucs
func (p *Person) SetName(name string) {
	p.Name = name
}

func methodMain() {
	pers := Person{"Vasily", 31}
	pers.SetName("Vasily Galperov") // = (&pers).SetName("Vasily Galperov")
	fmt.Printf("Updated person: %#v\n", pers)

	var acc Account = Account{
		ID:   1,
		Name: "VGalperov",
		Person: Person{
			Name: "Vasily",
			Age:  32,
		},
	}

	acc.SetName("Vasily.Galperov")
	fmt.Printf("Updated person: %#v\n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(3)
	fmt.Printf("MySlice count: %d\n", sl.Count())
}
