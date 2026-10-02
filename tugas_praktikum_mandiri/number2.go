package main

import "fmt"

func main() {
	var gp, wl, bl, pt, gb int64
	fmt.Scan(&gp, &wl)
	bl = wl * 45000
	pt = int64(float64(gp) * 0.055)
	gb = gp + bl - pt
	fmt.Println(gb)
}
