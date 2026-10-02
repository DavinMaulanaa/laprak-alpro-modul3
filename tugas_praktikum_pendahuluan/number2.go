package main

import "fmt"

func main (){
	var n int
	fmt.Scanln(&n)
	temp_position := n + 4
	if temp_position > 7{
		hitung_hari := (temp_position % 7)
		if hitung_hari == 0{
			hitung_hari = 7}
		fmt.Println(hitung_hari)
	
	} else {
		fmt.Println(temp_position)
	}

}