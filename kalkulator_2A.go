package main

import (
	"fmt"
	"os"	
	) 

var Pilih, Angka1, Angka2 int

func main(){		
	for {
		fmt.Println("Pilih Menu Perhitungan")
		fmt.Println("1. Penjumlahan 2 angka")	
		fmt.Println("2. Pengurangan 2 angka")
		fmt.Println("3. Perkalian 2 angka")
		fmt.Println("4. Pembagian 2 angka")
		fmt.Println("5. Keluar")
		fmt.Printf("Pilih : ") 
		fmt.Scan(&Pilih)
		if Pilih < 1 || Pilih > 5 {
			fmt.Println("Pilih Menu Hanya 1 sampai 5 !")
			fmt.Println()			
		}			
		switch Pilih {						
			case 1 : 
				fmt.Println("Penjumlahan 2 angka")
				InputAngka()
				fmt.Println("Hasil Penjumlahan ", Angka1, " + ",  Angka2, " = ", Angka1+Angka2)
				fmt.Println()
				break		
			case 2 :
				fmt.Println("Pengurangan 2 angka")		
				InputAngka()
				fmt.Println("Hasil Pengurangan ", Angka1, " - ", Angka2, " = ", Angka1-Angka2)										
				fmt.Println()		
				break
			case 3 :
				fmt.Println("Perkalian 2 angka")						
				InputAngka()
				fmt.Println("Hasil Perkalian ", Angka1, " x ", Angka2, " = ", Angka1*Angka2)										
				fmt.Println()
				break		
			case 4 :
				fmt.Println("Pembagian 2 angka")									
				InputAngka()
				fmt.Println("Hasil Pembagian ", Angka1, " / ", Angka2, " = ", Angka1/Angka2)										
				fmt.Println()	
				break
			case 5 :
				os.Exit(1)
		}
	}		
}

func InputAngka()(int, int){	
	fmt.Printf("Masukkan Angka ke 1 : ")
	fmt.Scan(&Angka1)
	fmt.Printf("Masukkan Angka ke 2 : ")
	fmt.Scan(&Angka2)
	return Angka1, Angka2		
}
