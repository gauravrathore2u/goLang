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
	//---------------------------------------------------------------------------------------------
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
	//---------------------------------------------------------------------------------------------

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
	//---------------------------------------------------------------------------------------------

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


	// redeclare to variable
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

	//Conditional statements
	//---------------------------------------------------------------------------------------------
	pee := 10;
	if pee == 10 {
		fmt.Println("pee is 10");
	} else if pee == 20 {
		fmt.Println("pee is 20");
	} else {
		fmt.Println("pee is not 10 or 20");
	}

	// the above thing can also done like
	//here we are declaring the variable inside the if block
	// this will make scope of lee variable limited to if block
	//eg. if height := getHeight(age); height < 5.0 { ... }
	if lee := 10; lee < 10 {   
		fmt.Println("lee is less than 10");
	} else if lee == 20 {
		fmt.Println("lee is 20");
	} else {
		fmt.Println("lee is not 10 or 20");
	}

	//Functions
	//---------------------------------------------------------------------------------------------
	/*
	func sub(a int, b int) int {
		return a - b;
	}
	*/

	fmt.Println(sub(10, 20));

	//for same type in arguments we can declare only to the last one
	/*
	func sub(a, b int) int {
		return a - b;
	}
	*/
	fmt.Println(sub1(10, 20));

	//we can return multiple value from a function
	/*
	func sub(a, b int) (int, int) {
		return a - b, a + b;
	}
	*/
	k, m := sub2(10, 20);
	fmt.Println(k, m);

	//we can ignore the return value using '_'
	/*
	func sub(a, b int) (int, int) {
		return a - b, a + b;
	}
	*/
	_, m = sub2(10, 20);
	fmt.Println(m);

	//Named return value
	//returned values may be given name, and if they are, then they are treated the same as 
	//if they were new variable defined at the top of the function 
	/*
	func getCord()(x, y int) {
		// here x and y are initialized with zero value
		
		return; //automatically return x and y
	}
	*/
	// the above code is same as
	/*
	func getCord1()(int, int) {
		var x int = 0;
		var y int = 0;
		return x, y;
	}
	*/
	//recommended is to declare the return variable but also return them



	
}

func sub(a int, b int) int {
	return a - b;
}
func sub1(a, b int) int {
	return a - b;
}
func sub2(a, b int) (int, int) {
	return a - b, a + b;
}
func getCord()(x, y int) {
	return;
}

func getCord1()(int, int) {
	var x int = 10;
	var y int = 20;
	return x, y;
}