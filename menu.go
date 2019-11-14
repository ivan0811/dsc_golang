package main

import (
	"fmt"
	"bufio"
	"os"
)

type User struct{
	NIM string
	Password string
	Nama string
}

var Pilih int;

var ListUser = []User{}

func main(){	
	var nim, password, nama string
	LengthList := 0;	
	scanner := bufio.NewReader(os.Stdin)						
	for{
		fmt.Println("Menu Pilihan")
		fmt.Println("1. Tambah data")
		fmt.Println("2. Cari Data")
		fmt.Println("3. Login")
		fmt.Println("4. Keluar")
		fmt.Scanln(&Pilih)		
		if Pilih < 1 || Pilih > 4 {
			fmt.Println("Pilih Menu Hanya 1 sampai 4 !")
		}
		switch Pilih {
		case 1 : 			
			fmt.Println("Masukkan NIM		:")
			nim, _= scanner.ReadString('\n')
			fmt.Println("Masukkan Password	:")
			password, _= scanner.ReadString('\n')
			fmt.Println("Masukkan Nama		:")
			nama, _= scanner.ReadString('\n')			
			ListUser = append(ListUser, User{nim, password, nama})
			LengthList = LengthList + 1;			
		case 2 : 						
			fmt.Println("Cari NIM")
			fmt.Println("Masukkan NIM	:")						
			nim, _= scanner.ReadString('\n')					
			HasilCari := false
			for i := 0; i < LengthList; i++ {
				if nim == ListUser[i].NIM{
					fmt.Println(ListUser[i].Nama)
					HasilCari = true					
				}
			}
			if HasilCari == false{
				fmt.Println("Data Tidak Ada")
			}
		case 3 :
			Login := false			
			indeks := 0
			for indeks < 3 && Login == false{
				fmt.Println("LOGIN")
				fmt.Println("Masukkan NIM 		: ")
				nim, _= scanner.ReadString('\n')			
				fmt.Println("Masukkan Password  : ")			
				password, _= scanner.ReadString('\n')			
				for i := 0; i < LengthList; i++{
					if nim == ListUser[i].NIM && password == ListUser[i].Password{
						Login = true
					}else{						
						Login = false						
					}
				}				
				if Login == false{
					fmt.Println("Username atau Password salah!")
					indeks++
				}							
			}			
			if Login == true{
				fmt.Println("Anda berhasil Login")
			}else if indeks <= 3 && Login == false{
				fmt.Println("anda terblokir")				
			}			
		case 4 :	
			os.Exit(1)
		}		
	}	
}
