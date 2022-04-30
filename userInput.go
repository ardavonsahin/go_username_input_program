package main

import "fmt"

func main(){
	var username  string
	var password  string
	var sys_username string = "arda"
	var sys_password  string = "dutt81"

	fmt.Println("Please Enter Your Username ID : ")
	fmt.Scanln(&username)

	fmt.Println("Please Enter Your Password : ")
	fmt.Scanln(&password)

	if username != sys_username && password == sys_password{
		fmt.Println("Invalid Username...")
	}else if username == sys_username && password != sys_password{
		fmt.Println("Invalid Password...")
	}else if username != sys_username && password != sys_password{
		fmt.Println("Invalid Username and Password...")
	}else{
		fmt.Println("Welcome, ",username)
	}
}