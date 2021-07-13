package main
import (
	"fmt"
)
func main(){
	fmt.Println("Hello, Black Hat Hacker")

	var x="Hola"
	fmt.Println(x)

	z := int(42)
	fmt.Println(z)


	fmt.Println("----SLICCE & MAP----")
	var s = make([]string, 0)
	var m= make(map[string]string)
	s= append(s, "some string")
	m["some key"] = "some value"


	fmt.Println("slcie: ")
	fmt.Println(s)
	fmt.Println("Map: ")
	fmt.Println(m)
	fmt.Println("----POINTERS----")
	var numero=int(45)
	pointer:= &numero
	fmt.Println("primer apuntador")
	fmt.Println(pointer)
	fmt.Println("retrive value stored: ")
	fmt.Println(*pointer)
	fmt.Println("cambiando el valor al numero..")
	*pointer=100
	fmt.Println(numero)

	fmt.Println("----ESTRUCTURAS___")


}
