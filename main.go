package main

import "fmt"

var a int

type hotdog int //Declaring my own type which is hotdog and what type of values can be assigned is int
var b hotdog //declaring b of type hotdog

func main()  {
	a=42;
	fmt.Println(a);
	fmt.Printf("%T\n",a)

	b=50;
	fmt.Println(b);
	fmt.Printf("%T\n",b)

	//a=b , this assignment is not feasible as a & b are two different type
	//but conversion is possible and below is the conversion
	a=int(b)
	fmt.Println(a)

}
