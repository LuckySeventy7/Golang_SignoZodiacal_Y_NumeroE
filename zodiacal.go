package main

import "fmt"

func main() {
	var dia int
	var mes int

	//fmt.Println("Ingrese, ")
	fmt.Println("Ingrese dia de nacimiento: ")
	fmt.Scan(&dia)
	

	fmt.Println("Ingrese mes de nacimiento: ")
	fmt.Scan(&mes)
	switch mes{
		case 1:{
			if dia >=21 && dia <=31{
				fmt.Println("Acuario")
			}else if dia <=20 && dia > 0{
				fmt.Println("Capricomio")
				}
			
		}
		case 2: {
			if dia >=20 && dia <=29{
				fmt.Println("Piscis")
			}else if dia <=19 && dia > 0{
				fmt.Println("Acuario")
			}
		}
		case 3: {
			if dia >=21 && dia <=31{
				fmt.Println("Aries")
			}else if dia <=20 && dia > 0{
				fmt.Println("Piscis")
			}
			
		}
		case 4:  {
			if dia >=21 && dia <=30{
				fmt.Println("Tauro")
			}else if dia <=20 && dia > 0{
				fmt.Println("Aries")
			}
		
		}
		case 5: {
			if dia >=21 && dia <=31{
				fmt.Println("Geminis")
			}else if dia <=20 && dia > 0{
				fmt.Println("Tauro")
			}
		
		}
		case 6: {
			if dia >=22 && dia <=30{
				fmt.Println("Cancer")
			}else if dia <=21 && dia > 0{
				fmt.Println("Gemisis")
			}
			
		}
		case 7: {
			if dia >=23 && dia <=31{
				fmt.Println("Leo")
			}else if dia <=22 && dia > 0{
				fmt.Println("Cancer")
			}
			
		}
		case 8: {
			if dia >=24 && dia <=30{
				fmt.Println("Virgo")
			}else if dia <=23 && dia > 0{
				fmt.Println("Leo")
			}
			
		}
		case 9:{
			if dia >=23 && dia <=30{
				fmt.Println("Libra")
			}else if dia <=22 && dia > 0{
				fmt.Println("Virgo")
			}
			
		}
		case 10: {
			if dia >=23 && dia <=31{
				fmt.Println("Escorpio")
			}else if dia <=22 && dia > 0{
				fmt.Println("Libra")
			}
			
		}
		case 11: {
			if dia >=23 && dia <=30{
				fmt.Println("Sagitario")
			}else if dia <=22 && dia > 0{
				fmt.Println("Escorpio")
			}
		
		}
		case 12: {
			if dia >=22 && dia <=31{
				fmt.Println("Capriomio")
			}else if dia <=21 && dia > 0{
				fmt.Println("Sagitario")
			}
		
		}
		
	}

}
