package demo2

import (
	"fmt"
	"unsafe"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	dogName string
}
type Cat struct {
}
type Fish struct {
}

func (d Dog) Speak() string {
	dog := &Dog{}
	dog1 := Dog{}
	fmt.Println(unsafe.Sizeof(dog))
	fmt.Println(unsafe.Sizeof(dog1))

	return "dog"
}
func (c Cat) Speak() string {

	return "cat"
}
func (f Fish) Speak() string {

	return "fish"
}
func main() {

	animals := []Animal{Dog{}, Cat{}, Fish{}}
	for _, animal := range animals {

		fmt.Println(animal.Speak())
	}
	d := &Dog{}
	fmt.Println(d.dogName)

}
