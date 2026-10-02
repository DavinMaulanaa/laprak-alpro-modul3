package main
import "fmt"
func main(){
// [DEKLARASI VARIABEL]
var kecepatan, menit, jam, total_jarak int //MASUKKAN TIPE DATA
// [INPUT]
fmt.Scan(&kecepatan) // MASUKKAN VARIABEL INPUT
// [PROSES]
total_jarak = 100+60+170 //MASUKKAN TOTAL JARAK TEMPUH (100+60+170)
menit = (total_jarak * 60) / kecepatan //BERISI TOTAL JARAK
jam = menit/60 //MASUKKAN RUMUS MENCARI JAM
menit = menit % 60 //MASUKKAN SISA MENIT
// [OUTPUT]
fmt.Println(jam, menit) //MASUKKAN VARIABEL YANG TELAH DIPROSES
}