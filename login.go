package main

import "fmt"

func main(){
	var NIM, Pass string
	var login bool
	login = false	
	var IndexNim = [2]string{"1011", "1012"}
	var IndexPass = [2]string{"0811", "1234"}	
	i := 0	
	for i < 3 && login == false{
		fmt.Println("Masukan NIM : ")
		fmt.Scan(&NIM)
		fmt.Println("Masukan Password : ")
		fmt.Scan(&Pass)	
		for index := 0; index < len(IndexNim); index++{
			if NIM == IndexNim[index] && Pass == IndexPass[index]{
				login = true
				fmt.Println("Anda Berhasil Login")				
			}
		}	
		if login == false {
			fmt.Println("Maaf NIM dan Password Salah silahkan coba lagi")			
			i++									
		}		
	}	
	if i == 3 && login == false {
		fmt.Println("Maaf anda terblokir")		
	} 
}
