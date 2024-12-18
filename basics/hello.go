//go is a compiled language
//go is a statically typed language
//go is a garbage collected language
//go is a concurrent language
//go compiles faster than java

//to run go file: go run hello.go

package main

import (
	"errors"
	"fmt"
)


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



	//Struct
	//---------------------------------------------------------------------------------------------
	//struct is similar to the object in javascript and dictionary in python
	type Address struct {
		City string;  // all values are assigned initial values as "" or 0 or false
		State string;
	}
	type Person struct {
		Name string;
		Age int;
		IsStudent bool;
		PersonAddress Address;
	}
	
	p2  := Person{}; // initialize the struct
	var p1 Person = Person{"John", 30, true, Address{"New York", "NY"}};
	fmt.Println(p1);
	fmt.Println((p2.IsStudent));

	//Embedded struct
	type Person1 struct {
		Address;  // embedded struct, now all the properties of Address will be available
		Name string;
		Age int;
	}
	type Address1 struct {
		City string;
		State string;
	}
	p3 := Person1{};
	p4 := Person1{Address{"New York", "NY"}, "John", 30};
	fmt.Println(p3);
	fmt.Println(p4.City); // now we can directly access the property of Address


	//Struct methods
	//---------------------------------------------------------------------------------------------
	p5 := rect{10, 20};
	fmt.Println(p5.area());



	//Interface
	//---------------------------------------------------------------------------------------------
	/*
	an interface is a type that lists methods without providing their code. 
	You can’t create an instance of an interface directly, but you can make a variable of the interface 
	type to store any value that has the needed methods.
	A type implements an interface by by implementing its methods. unlike in other languages,
	there is no explicit declaration of intent, there is no 'implements' keyword.

	type InterfaceName interface {
    	Method1() returnType
   		Method2() returnType
	}
	*/

	var s Shape;
	s = Rectangle{length: 4, width: 3}
    fmt.Println("R Area:", s.Area())
    fmt.Println("R Perimeter:", s.Perimeter())


	//Errors
	//---------------------------------------------------------------------------------------------
	//in js we handle errors using try catch
	//in go we handle errors using error interface, errors.New("error message")
	//we can return error with return type and check errors where needed

	val, err := divideNum(10, 0);
	if err != nil {
		fmt.Println(err);
		// return;  //we can return here
	} else {
		fmt.Println(val);
	}


	//Loops
	//---------------------------------------------------------------------------------------------
	/*
		for INITIALIZATION; CONDITION; AFTER {}
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}
	//Similar to other languages, we can break the loop using break keyword
	//we can continue the loop using continue keyword


	//infiinite loop
	/*
	if we don't have any condition then the loop will run forever
	for INITIALIZATION; ; AFTER {}
	*/

	//While loop
	//in go we don't have while loop, we can use for loop with a condition
	/*
	for CONDITION {}
	*/
	plantHeight := 1;
	for plantHeight < 10 {
		fmt.Println("Plant height:", plantHeight)
		plantHeight++
	}


	//Arrays
	//---------------------------------------------------------------------------------------------
	/*
		arrayName := [5]int{1, 2, 3, 4, 5}
		[lengthOfArray] datatype {values}
	*/
	var nummies = [5] int {1, 2};  // length of array is 5 and initialized with 0
	fmt.Println(nummies);		// [1 2 0 0 0]

	//Infer length of array
	//we can also declare array without length
	var nummies1 = [2] int {1, 2}; // length of array is 2
	fmt.Println(nummies1);		// [1 2]

	//assign values to certain positions 
	 nummies2 := [5] int {2: 1, 4: 2} // length of array is 5; and at index 2 value is 1 and at index 4 value is 2
	 fmt.Println(nummies2);		// [0 0 1 0 2]

	 //Multidimensional array
	 nummies3 := [2][2]int{{1, 2}, {3, 4}}
	 fmt.Println(nummies3);		// [[1 2] [3 4]]


	 //Slice
	 //---------------------------------------------------------------------------------------------
	 /*
	 slices are similar to array but we don't provide length, it has dynamic length.
	 we can add or remove elements form slice.
	 sliceName := []datatype{values}  //in slices we do not provide length
	 */
	 var nummies4 = []int{1, 2, 3}
	 nummies4 = append(nummies4, 5);
	 fmt.Println(nummies4);		// [1 2 3 5]

	 //we can extract range from the slice
	 fmt.Println(nummies4[1:3])   //exclusive of index 3
	 fmt.Println(nummies4[:3])    //exclusive of index 3
	 fmt.Println(nummies4[1:])   //inclusive of index 3

	 //Slice using 'make'
	nummies5 := make([]int, 2);  //create a slice of length 2
	nummies5[0] = 1;
	nummies5[1] = 2;
	// nummies5[2] = 3; //this will throw error

	//but we can use all the slice methods
	fmt.Println(len(nummies5));  //2
	fmt.Println(append(nummies5, 6));       // [1 2 6]

	//append a slice to a slice
	nummies6 := []int{1, 2, 3}
	nummies7 := []int{4, 5}
	nummies6 = append(nummies6, nummies7...)  //append a slice to a slice
	fmt.Println(nummies6)
	fmt.Println(len(nummies6))


	//Map
	//---------------------------------------------------------------------------------------------
	/*
	 mapName := map[keyType]valueType{key: value}
	*/
	var nummies8 = map[string]int{"one": 1, "two": 2}
	fmt.Println(nummies8)
	fmt.Println(nummies8["one"]);

	//map using make 
	nummies9 := make(map[string]int);
	nummies9["one"] = 1;
	nummies9["two"] = 2;
	fmt.Println(nummies9) 
	fmt.Println(len(nummies9)) 

	//get value from map
	var oneVal = nummies9["one"]; 
	fmt.Println(oneVal);

	//check if value exists
	val, ok := nummies9["one"];
	if ok { //if value exists then ok will be true
		fmt.Println(val);
	}

	//delete value from map
	delete(nummies9, "one");
	fmt.Println(nummies9) 
 
	//Imp:- if we try to get the element which is not present in the map then it will return zero value

}

type rect struct {
	width  int
	height int
	// newArea func() //we can also give methods to struct but this will loose the 
					   //singularity we separatly has to save the area in function , 
					// if someone change the width in struct and forgot to change the area then it will be wrong
					// we can also give methods to struct by declaring struct before fn name
}
// we can give methods to struct by declaring struct before fn name
func (r rect) area() int {
	return r.width * r.height
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

//Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle type that implements the Shape interface
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func divideNum(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}