package main

import "fmt"

//type Assertion allows extracting the underlying type of an interface like in this code it find the value of s.(Shape)
//Syntax: value := interfaceVariable.(ConcreteType)
func printArea(s interface{}) {
	value, ok := s.(Shape)
	if ok {
		fmt.Println("Area:", value.Area())
	} else {
		fmt.Println("Not a Shape interface")
	}
}

type Shape interface {
	Area() float64
}
type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}
func main() {
	sq := Square{side: 4}
	printArea(sq)
}
