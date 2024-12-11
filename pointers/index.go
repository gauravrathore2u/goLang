package main

import "fmt";

func main() {
	var a int = 10;
	var b *int = &a;  //b is a pointer to address of a
	fmt.Println(a); //10
	fmt.Println(b); //address of a
	fmt.Println(*b); //dereference //value at address of a i.e. 10

	//Imp:- Pointers can be dangerous
	//If a pointer points to nothing (zero value of the pointer type ) then dereferencing it will cause a run time error.
	//that crashes the program.
	//whenever dealing with pointers always check if the pointer is nil before dereferencing it.
}
