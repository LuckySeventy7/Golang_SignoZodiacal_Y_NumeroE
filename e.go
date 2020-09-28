package main

import "fmt"

func main() {
	i:=1
	e:=  2.71828;
	var limite int;
	result:=e;
	fmt.Print("calcular e^x \n")
	fmt.Printf("Ingrese limite (numero entero): ")
	fmt.Scan(&limite);
	for i < limite{
		result = result * e; 
		i++

	}
	fmt.Printf("respuesta %f", result);

}