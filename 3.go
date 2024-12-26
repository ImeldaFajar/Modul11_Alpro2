package main

import "fmt"

const NMAX = 100000

var data [NMAX]int

func main() {
	var n, k int

	// Membaca input n dan k
	fmt.Scan(&n, &k)

	// Mengisi array dengan n bilangan
	isiArray(n)

	// Mencari posisi k dalam array
	pos := posisi(n, k)

	// Output hasil
	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

func isiArray(n int) {
	/* I.S. Array data terdefinisi hingga indeks n-1 */
	/* F.S. Array diisi dengan n bilangan */
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	/* Mencari posisi k dalam array data hingga indeks n-1 */
	/* Posisi dimulai dari 0. Jika tidak ditemukan, kembalikan -1 */
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}
	return -1
}
