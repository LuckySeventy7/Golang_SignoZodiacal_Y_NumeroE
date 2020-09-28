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
	if  mes == 3 {
		if dia >=21 && dia <=31{
			fmt.Println("Aries")
		}else if dia <=20 && dia > 0{
			fmt.Println("Piscis")
		}else {
			fmt.Println("error ")
		}
	}else if mes == 4 {
		if dia >=21 && dia <=30{
			fmt.Println("Tauro")
		}else if dia <=20 && dia > 0{
			fmt.Println("Aries")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 5 {
		if dia >=21 && dia <=31{
			fmt.Println("Geminis")
		}else if dia <=20 && dia > 0{
			fmt.Println("Tauro")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 6 {
		if dia >=22 && dia <=30{
			fmt.Println("Cancer")
		}else if dia <=21 && dia > 0{
			fmt.Println("Gemisis")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 7 {
		if dia >=23 && dia <=31{
			fmt.Println("Leo")
		}else if dia <=22 && dia > 0{
			fmt.Println("Cancer")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 8 {
		if dia >=24 && dia <=30{
			fmt.Println("Virgo")
		}else if dia <=23 && dia > 0{
			fmt.Println("Leo")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 9 {
		if dia >=23 && dia <=30{
			fmt.Println("Libra")
		}else if dia <=22 && dia > 0{
			fmt.Println("Virgo")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 10 {
		if dia >=23 && dia <=31{
			fmt.Println("Escorpio")
		}else if dia <=22 && dia > 0{
			fmt.Println("Libra")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 11 {
		if dia >=23 && dia <=30{
			fmt.Println("Sagitario")
		}else if dia <=22 && dia > 0{
			fmt.Println("Escorpio")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 12 {
		if dia >=22 && dia <=31{
			fmt.Println("Capriomio")
		}else if dia <=21 && dia > 0{
			fmt.Println("Sagitario")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 1 {
		if dia >=21 && dia <=31{
			fmt.Println("Acuario")
		}else if dia <=20 && dia > 0{
			fmt.Println("Capricomio")
		}else {
			fmt.Println("error ")
		}
	}else if  mes == 2 {
		if dia >=20 && dia <=29{
			fmt.Println("Piscis")
		}else if dia <=19 && dia > 0{
			fmt.Println("Acuario")
		}else {
			fmt.Println("error ")
		}
	}else{
		fmt.Println("error ")
	}

}
