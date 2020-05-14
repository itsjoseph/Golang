package main

import "fmt"

func salidas()  {
	count := 501
	for i := 0; i < count; i++ {
		fmt.Println("Saliendo del metodo: ",i)
	}
}

func main(){
	count := 501	
	for i := 0; i < count; i++ {
		salidas()
		fmt.Println(i)
	}
	fmt.Println("Hola como estas")
}