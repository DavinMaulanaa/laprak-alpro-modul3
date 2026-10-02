package main

import "fmt"

func main(){
	var koin, hitung int
	fmt.Scan(&koin)
	hitung = (koin)
	emas := hitung / 9
	sisa := hitung % 9
	perak := sisa / 3
	tembaga := sisa % 3
	fmt.Println(emas, perak, tembaga)
}