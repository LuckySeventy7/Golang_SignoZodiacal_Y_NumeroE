package main

import "fmt"

func main() {
	var n float64
	var j float64
	var i float64
	var fact float64//factorial
	var e float64

	e=1; // se inicia con 1 para tomar en cuenta el  0!
	fmt.Scan(&n)
	//e=0; //rep factorial 0
	for j = n; j > 0; j-- {
		fact = 1 
		for i = 1; i <= j;i++ {
			fact = fact * i;		
		}
		e = e + 1.0/fact;
	}
	fmt.Println(e)
}