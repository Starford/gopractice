package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 90.33
	d := true

	//but can do the declare and assign
	var g = "ule mse"

	//i can also declare empty variables
	var emptyint int
	var emptystring string
	var emptyfloat float64
	var emptybool bool

	//can also do fancy things like
	var male, female string = "male", "female"

	/* but this is considered the best i dont know why */
	var animal bool         //zero values
	dinosour := "t-rex"     //declare and initialize
	cow, goat := true, true //declare and initialize one or more example

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Printf("%T \n", g)

	fmt.Printf("%v \n", emptyint)
	fmt.Printf("%v \n", emptystring)
	fmt.Printf("%v \n", emptyfloat)
	fmt.Printf("%v \n", emptybool)

	fmt.Printf("%v \n", male)
	fmt.Printf("%v \n", female)

	fmt.Printf("%v \n", cow)
	fmt.Printf("%v \n", goat)
	fmt.Printf("%v \n", animal)
	fmt.Printf("%v \n", dinosour)
}
