package main

import "fmt"

func main() {
	var dia int
	var mes int

	//fmt.Println("Ingrese, ")

	fmt.Scan(&dia)

	fmt.Scan(&mes)
	switch mes{
		case 1:{
			if dia >=21 && dia <=31{
				fmt.Println("acuario")
			}else if dia <=20 && dia > 0{
				fmt.Println("capricomio")
				}
			
		}
		case 2: {
			if dia >=20 && dia <=29{
				fmt.Println("piscis")
			}else if dia <=19 && dia > 0{
				fmt.Println("acuario")
			}
		}
		case 3: {
			if dia >=21 && dia <=31{
				fmt.Println("aries")
			}else if dia <=20 && dia > 0{
				fmt.Println("piscis")
			}
			
		}
		case 4:  {
			if dia >=21 && dia <=30{
				fmt.Println("tauro")
			}else if dia <=20 && dia > 0{
				fmt.Println("aries")
			}
		
		}
		case 5: {
			if dia >=21 && dia <=31{
				fmt.Println("geminis")
			}else if dia <=20 && dia > 0{
				fmt.Println("tauro")
			}
		
		}
		case 6: {
			if dia >=22 && dia <=30{
				fmt.Println("cancer")
			}else if dia <=21 && dia > 0{
				fmt.Println("geminis")
			}
			
		}
		case 7: {
			if dia >=23 && dia <=31{
				fmt.Println("leo")
			}else if dia <=22 && dia > 0{
				fmt.Println("cancer")
			}
			
		}
		case 8: {
			if dia >=24 && dia <=30{
				fmt.Println("virgo")
			}else if dia <=23 && dia > 0{
				fmt.Println("Leo")
			}
			
		}
		case 9:{
			if dia >=23 && dia <=30{
				fmt.Println("libra")
			}else if dia <=22 && dia > 0{
				fmt.Println("virgo")
			}
			
		}
		case 10: {
			if dia >=23 && dia <=31{
				fmt.Println("escorpio")
			}else if dia <=22 && dia > 0{
				fmt.Println("libra")
			}
			
		}
		case 11: {
			if dia >=23 && dia <=30{
				fmt.Println("sagitario")
			}else if dia <=22 && dia > 0{
				fmt.Println("escorpio")
			}
		
		}
		case 12: {
			if dia >=22 && dia <=31{
				fmt.Println("capriomio")
			}else if dia <=21 && dia > 0{
				fmt.Println("sagitario")
			}
		
		}
		
	}

}
