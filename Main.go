package main

import "fmt"

//	func main(){
//		var i int
//		i =42
//		var j int = 27
//		k := 99
//		fmt.Printf("%v , %T",j,j)
//	}
var (
	actorName    string = "Elisabeth Saden"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

// func main() {
// 	newman := "Hello"
// 	fmt.Printf("%v, %T",season,newman)
// 	var theURL string = "http://google.com"
// 	fmt.Println(theURL)
// }

//! variable type conversions

func main() {
	// var i int = 42
	// fmt.Printf("%v, %T\n", i,i)

	// var j string
	// j = strconv.Itoa(i)
	// fmt.Printf("%v, %T\n", j,j)

	// var n bool

	// var n uint16 = 42
	// fmt.Printf("%v, %T\n", n,n)

	// a := 10 //1010
	// b := 3  //0011
	// fmt.Println(a+b)
	// fmt.Println(a-b)
	// fmt.Println(a*b)
	// fmt.Println(a/b)
	// fmt.Println(a%b)

	// //? 2 int's of diffrent kinds can not be operated upon

	// fmt.Println(a&b)  //AND //0010
	// fmt.Println(a|b)  //OR //1011
	// fmt.Println(a^b)  //XOR //1001
	// fmt.Println(a&^b) //0100

	//* floating point numbers
	// var n float64 = 3.14
	// n = 13.7e72
	// n =2.1E14
	// fmt.Printf("%v, %T\n",n,n)

	//! complex numbers
	//var n complex128 = 1+2i
	// var n complex128 = complex(5,12)
	// fmt.Printf("%v, %T\n", real(n),real(n))
	// fmt.Printf("%v, %T\n", imag(n),imag(n))
	// fmt.Printf("%v, %T\n", n,n)

	// s := "this is a string"
	// s2 := "this is also a string"
	// b := []byte(s)
	// fmt.Printf("%v, %T\n", b, b)

	//* rune
	// var r rune = 'a'
	// fmt.Printf("%v, %T\n", r, r)

	//const myConst int = 42
	// myConst = 21 // ERROR
	//const myConst float64 = math.Sin(90)
	//fmt.Printf("%v, %T\n", myConst, myConst)

	//! constants

	// const (
	// 	a = iota
	// 	b
	// 	c
	// )

	// const (
	// 	a2 = iota
	// )

	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", a2)

	//* enumerated constants
	//? can be bit shifted
	const (
		_ = iota + 5
		catSpecialist
		dogSpecialist
		snakeSpecialist
	)

	var specialistType int = catSpecialist
	fmt.Printf("%v->%v\n", specialistType == catSpecialist, catSpecialist)
	fmt.Printf("%v\n", specialistType == dogSpecialist)
}
