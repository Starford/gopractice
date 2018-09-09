package main

import "fmt"

func main() {
	//use of functions
	greet("star")
	greet("ford")

	//can also do this, anonymous function
	greetings := func() {
		fmt.Println("called inside main")
	}
	greetings()

	//returns
	fmt.Println(salimiana("hello", "world"))

	//named return values
	fmt.Println(salamz("hello", "world"))

	//when you have more than one return
	fmt.Println(rudishamob("hello", "world"))

	//variadic functions which means it accepts n number of params
	fmt.Println(average(1, 3, 4, 5, 6, 5))
	/* but based on the above something like this can be done */
	floatnumbers := []float64{23, 15, 1, 45, 56, 23, 10, 3, 3, 9}
	fmt.Println(average(floatnumbers...))

	//return a function
	hi := makeGreeter("star")
	fmt.Println(hi())
	fmt.Printf("%T \n", hi)

}

func greet(name string) {
	fmt.Println(name)
}

func salimiana(firstname string, lastname string) string {
	return fmt.Sprint(firstname, lastname)
}

func salamz(firstname string, lastname string) (s string) {
	s = fmt.Sprint(firstname, lastname)
	return
}

func rudishamob(firstname string, lastname string) (string, string) {

	return fmt.Sprint(firstname, lastname), fmt.Sprint(lastname, firstname)
}

func average(st ...float64) float64 {

	var total float64
	//loop through
	for _, v := range st {
		total += v
	}
	return total / float64(len(st))
}

func makeGreeter(trailend string) func() string {
	return func() string {
		return "hello world from" + trailend
	}
}
