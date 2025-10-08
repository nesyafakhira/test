package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial tanpa rekursi
func faktorial(n int) int {
	result := 1                  // Inisialisasi hasil faktorial dengan 1
	for i := 2; i <= n; i++ {    // Loop dari 2 hingga n
		result *= i             // Mengalikan result dengan nilai i saat ini
	}
	return result                // Mengembalikan hasil faktorial
}

// Fungsi untuk menghitung permutasi
func permutasi(x, y int) int {
	if x > y {                   // Jika x lebih besar dari y
		return faktorial(x) / faktorial(x-y)  // Hitung x permutasi y: x! / (x-y)!
	}
	return faktorial(y) / faktorial(y-x)  // Jika y lebih besar dari x, hitung y permutasi x
}

func main() {
	var x, y int                 // Deklarasi variabel untuk input user

	// Meminta input dari user
	fmt.Print("Masukkan dua bilangan bulat: ")
	fmt.Scan(&x, &y)              // Membaca input x dan y dari user

	// Menghitung faktorial x dan y
	xFact := faktorial(x)
	yFact := faktorial(y)

	// Menghitung nilai permutasi berdasarkan nilai x dan y
	hasilPermutasi := permutasi(x, y)

	// Menampilkan hasil faktorial dan permutasi
	fmt.Println(xFact, yFact, hasilPermutasi)
}
