package main

import (
	"fmt"

	"github.com/fatih/color"
)

func Payment() {
	Harga := 0
	fmt.Println("Harga Barang: ")
	fmt.Scan(&Harga)

	Uang := 0
	fmt.Println("Uang Pembeli: ")
	fmt.Scan(&Uang)

	if Uang < Harga {
		kurang := Harga - Uang
		color.Red("[SISTEM]: Transaksi Ditolak, Uang kurang %d\n", kurang)
	} else if Harga == Uang {
		color.Green("[SISTEM]: Transaksi Berhasil, Uang Anda Pas")
	} else {
		kembalian := Uang - Harga
		color.Green("[SISTEM]: Transaksi Berhasil, Kembalian %d\n", kembalian)
	}
}
func main() {
	Payment()
}
