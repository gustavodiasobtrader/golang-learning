package main

import "fmt"

func main(){
  var num int
  fmt.Println("Digite um número de 1 a 12 verá a qual mês corresponde.")
  fmt.Scan(&num)

  switch num {
  	case 1:
  		fmt.Println("Janeiro")
  	case 2:
  		fmt.Println("Fevereiro")
  	case 3:
  		fmt.Println("Março")
  	case 4:
  		fmt.Println("Abril")
  	case 5:
  		fmt.Println("Maio")
  	case 6:
  		fmt.Println("Junho")

  	case 7:
  		fmt.Println("Julho")
  	case 8:
  		fmt.Println("Agosto")
  	case 9:
  		fmt.Println("Setembro")
  	case 10:
  		fmt.Println("Outubro")
  	case 11:
  		fmt.Println("Novembro")
  	case 12:
  		fmt.Println("Dezembro")
  	default:
  		fmt.Println("\nMês inexistente. Favor digitar um valor válido. \n\n")
  		main()
  }


}
