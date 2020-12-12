package main

import "fmt"
import "strings"
// import "math/rand"
// import "time"

// bisa dengan
// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

func main()  {
	// var nama string = "Bayu"
	// penulisan variabel tanpa menulis var dan tipe datanya
	// variabel yang tipe datanya ditentukan oleh nilainya
	// name = "Bayu"

	// pointer
	// name := new(string)
	// fmt.Println(name)   // 0x20818a220 alamatnya
	// fmt.Println(*name)  // "" nilai datanya

	// var kota string
	// kota = "Yogyakarta"

	// deklarasi multi variabel
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// dengan type inference cara singkat dapat ditulis walaupun tipe data berbeda
	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// variabel _ digunakan untuk variabel yang belum digunakan
	// _ = "predefined"

	// print line dengan variabel
	// fmt.Printf("halo %s yang tinggal di %s!\n", nama, kota)
	// print line tanpa variabel
	// fmt.Println("Hello World")
	// penggabungan string menggunakan tanda +
	// fmt.Println("halo", nama, kota + "!")

	// deklarasi tipe data non desimal
	// var positiveNumber uint8 = 89
	// var negativeNumber = -1243423644

	// %d digunakan untuk tipe data non desimal
	// fmt.Printf("bilangan positif: %d\n", positiveNumber)
	// fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// deklarasi tipe data desimal
	// var decimalNumber = 2.62

	// %f digunakan untuk bilangan desimal
	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// deklarasi tipe data boolean
	// var exist bool = true
	// %t digunakan untuk data boolean
	// fmt.Printf("exist? %t \n", exist)

	// konstanta variabel
	// const firstName string = "john"
	// fmt.Print("halo ", firstName, "!\n")

	// type inferencenya
	// const lastName = "wick"
	// fmt.Print("nice to meet you ", lastName, "!\n")

	// pengondisian
	// var point = 8

	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulus")
	// } else {
	// 	fmt.Printf("tidak lulus. nilai anda %d\n", point)
	// }

	// variabel temp pada if else
	// var point = 8840.0

	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }

	// switch case
	// var point = 6

	// switch point {
	// case point == 8: // swtich harus switch {} aja
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// 	fallthrough // akan dilanjutkan pengecekan pada case selanjutnya tanpa menghiraukan nilai kondisinya
	// default:
	// 	{ // bisa dengan {}
	// 		fmt.Println("not bad")
	// 	}
	// }

	// perulangan for
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i)
	// }
	// var i = 0

	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}

	// 	fmt.Println()
	// }

	// outerLoop: // label untuk for dibawahnya
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		if i == 3 {
	// 			break outerLoop // berhentikan perulangan dengan label outerloop
	// 		}
	// 		fmt.Print("matriks [", i, "][", j, "]", "\n")
	// 	}
	// }

	// array
	// array ditentukan jumlahnya yaitu 4
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}


	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)

	// array tidak ditentukan jumlahnya
	// var	numbers = [...]int{2, 3, 2, 4, 3}

	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))

	// array multidimensional
	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// var fruits = make([]string, 2)
	// fruits[0] = "apple"
	// fruits[1] = "manggo"

	// fmt.Println(fruits)  // [apple manggo]

	// slice
	// perbedaan dengan array yaitu jumlah tidak perlu didefinisikan jumlahnya
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fungsiPertama("tes", fruits)
	// fmt.Println(fruits[0]) // "apple"

	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// var newFruits = fruits[0:2]

	// fmt.Println(newFruits) // ["apple", "grape"]

	// len(fruits) menghitung panjang jumlah elemen slice fruits

	// cap(fruits) menghitung lebar atau kapasitas maksimum slice yang sama seperti slice aslinya

	// var cFruits = append(fruits, "papaya") // append digunakan untuk menambah nilai pada slice

	// copy(destination, source) // digunakan untuk mengcopy elemen slice

	// map
	// asosiatif pada go
	// var chicken map[string]int
	// chicken = map[string]int{}

	// chicken["januari"] = 50
	// chicken["februari"] = 40

	// fmt.Println("januari", chicken["januari"]) // januari 50
	// fmt.Println("mei",     chicken["mei"])     // mei 0

	// var chicken = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }

	// for key, val := range chicken {
	// 	fmt.Println(key, "  \t:", val)
	// }

	// delete(chicken, "januari") // digunakan untuk menghapus niali map berdasarkan keynya

	// check apakah key ada
	// var chicken = map[string]int{"januari": 50, "februari": 40}
	// var value, isExist = chicken["mei"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }
	// var chickens = []map[string]string{
	// 	{"name": "chicken blue",   "gender": "male"},
	// 	{"name": "chicken red",    "gender": "male"},
	// 	{"name": "chicken yellow", "gender": "female"},
	// }

	// rand.Seed(time.Now().Unix()) // generate angka random

	// fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) // %v digunakan untuk menampilkan segala jenis data
}

func fungsiPertama(message string, arr []string){
	var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}

// jika ingin menggunakan return
// func randomWithRange(min, max int) int
