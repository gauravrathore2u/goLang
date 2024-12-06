//go is a compiled language
//go is a statically typed language
//go is a garbage collected language
//go is a concurrent language
//go compiles faster than java

//to run go file: go run hello.go

package main

import "fmt";

func main() {
	fmt.Println("Hello World!");

	//variables
	/*
	//1. using data type declaration ==> var variableName type = value
	var x int = 10;
	var y float32 = 3.14; //32 bit float
	var y float64 = 3.14; //64 bit float
	var z bool = true;
	var s string = "Hello World";

	//2. using type inferred ==> variableName := value
	x := 10;
	y := 3.14;
	z := true;
	s := "Hello World";
	*/

	name  := "John";
	age  := 30;
	isStudent := true;
	fmt.Println("Name:", name);
	fmt.Printf("Name: %v", name);  // v is for value, use when not sure what is the data type
	fmt.Println("Name:", name);
	fmt.Printf("Name: %s", name);
	fmt.Println("Age:", age);
	fmt.Printf("Age: %d", age);
	fmt.Println("Is Student:", isStudent);
	fmt.Printf("Is Student: %t", isStudent);



	//Sprintf returns the formatted string
	var message string = fmt.Sprintf("Name: %s, Age: %d, Is Student: %t", name, age, isStudent);
	fmt.Println(message);


	//const keyword
	// const myName string = "John";
	// myName = "doe";   	//==>error
	// fmt.Println(myName);

	// const myName1 string;
	// myName1 = "doe";		//==>error
	// fmt.Println(myName1);


	/*
	Use := for local variables inside functions, as it is concise and idiomatic.
	Use var for package-level declarations, or when explicit type declarations are necessary.
	*/

	//multi variables declaration
	var (
		a2 int = 10;
		b2 int = 20;
		c2 int = 30;
	);

	var a1, b1, c1 int = 10, 20, 30;
	a, b, c := 10, 20, 30;
	fmt.Println(a2, b2, c2);
	fmt.Println(a1, b1, c1);
	fmt.Println(a, b, c);


	// return to variable
	var p int = 10;
	p = 20;
	// p = "name"; //error
	fmt.Println(p);
	var q int;  //initialized as 0
	var t string; //initialized as ""
	fmt.Println(q, t);
	q = 10;
	t = "name";
	fmt.Println(q, t);
}