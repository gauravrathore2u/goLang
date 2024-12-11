package main

import "fmt";

func main() {
	var a int = 10;
	var b *int = &a;  //b is a pointer to address of a
	fmt.Println(a); //10
	fmt.Println(b); //address of a
	fmt.Println(*b); //value at address of a i.e. 10
}
