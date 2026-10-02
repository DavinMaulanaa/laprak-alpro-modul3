package main

import "fmt"

func main(){
	var n int
	var detik, menit, jam int
	fmt.Scanln(&n)
	detik = n % 60
	menit = (n%3600) / 60
	jam = n / 3600
	fmt.Println(jam, menit, detik)
} 