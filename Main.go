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

// func main() {
// 	// var i int = 42
// 	// fmt.Printf("%v, %T\n", i,i)

// 	// var j string
// 	// j = strconv.Itoa(i)
// 	// fmt.Printf("%v, %T\n", j,j)

// 	// var n bool

// 	// var n uint16 = 42
// 	// fmt.Printf("%v, %T\n", n,n)

// 	// a := 10 //1010
// 	// b := 3  //0011
// 	// fmt.Println(a+b)
// 	// fmt.Println(a-b)
// 	// fmt.Println(a*b)
// 	// fmt.Println(a/b)
// 	// fmt.Println(a%b)

// 	// //? 2 int's of diffrent kinds can not be operated upon

// 	// fmt.Println(a&b)  //AND //0010
// 	// fmt.Println(a|b)  //OR //1011
// 	// fmt.Println(a^b)  //XOR //1001
// 	// fmt.Println(a&^b) //0100

// 	//* floating point numbers
// 	// var n float64 = 3.14
// 	// n = 13.7e72
// 	// n =2.1E14
// 	// fmt.Printf("%v, %T\n",n,n)

// 	//! complex numbers
// 	//var n complex128 = 1+2i
// 	// var n complex128 = complex(5,12)
// 	// fmt.Printf("%v, %T\n", real(n),real(n))
// 	// fmt.Printf("%v, %T\n", imag(n),imag(n))
// 	// fmt.Printf("%v, %T\n", n,n)

// 	// s := "this is a string"
// 	// s2 := "this is also a string"
// 	// b := []byte(s)
// 	// fmt.Printf("%v, %T\n", b, b)

// 	//* rune
// 	// var r rune = 'a'
// 	// fmt.Printf("%v, %T\n", r, r)

// 	//const myConst int = 42
// 	// myConst = 21 // ERROR
// 	//const myConst float64 = math.Sin(90)
// 	//fmt.Printf("%v, %T\n", myConst, myConst)

// 	//! constants

// 	// const (
// 	// 	a = iota
// 	// 	b
// 	// 	c
// 	// )

// 	// const (
// 	// 	a2 = iota
// 	// )

// 	// fmt.Printf("%v\n", a)
// 	// fmt.Printf("%v\n", b)
// 	// fmt.Printf("%v\n", c)
// 	// fmt.Printf("%v\n", a2)

// 	//* enumerated constants
// 	//? can be bit shifted
// 	const (
// 		_ = iota + 5
// 		catSpecialist
// 		dogSpecialist
// 		snakeSpecialist
// 	)

// 	var specialistType int = catSpecialist
// 	fmt.Printf("%v->%v\n", specialistType == catSpecialist, catSpecialist)
// 	fmt.Printf("%v\n", specialistType == dogSpecialist)

// 	//* bit shiftings on enum cosnts
// 	const (
// 		_  = iota // ignore the first value by assigning a blank identifier
// 		KB = 1 << (10 * iota)
// 		MB
// 		GB
// 		TB
// 		PB
// 		EB
// 		ZB
// 		YB
// 	)

// 	fileSize := 4000000000.
// 	fmt.Printf("%.2f GB\n", fileSize/GB)

// 	const (
// 		isAdmin = 1 << iota
// 		isHeadquaters
// 		canSeeFinancials

// 		canSeeAfrica
// 		canSeeAsia
// 		canSeeEurope
// 		canSeeNorthAmerica
// 		canSeeSouthAmerica
// 	)

// 	var role byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", role)
// 	fmt.Printf("Is Admin? %v\n", isAdmin&role == isAdmin)
// 	fmt.Printf("Is HQ? %v\n", isAdmin&role == isHeadquaters)

// 	//!PascalCase for exported constants
// 	//!camelCase for internal constants
// }

//?Arrays
func main() {
	grades := [...]int{97, 85, 93}
	var students [3]string
	students[0] = "Lisa"
	students[1] = "Asad"
	students[2] = "GG"
	fmt.Printf("Grades %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Len(students): %v\n", len(students))

	//* 2d array
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	//* copy array
	// a := []int{1, 2, 3}
	// // b := a
	// // b[1] = 5
	// // c := &a
	// // c[1] = 5
	// fmt.Println(a)
	// // fmt.Println(b)
	// // fmt.Println(c)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	//* slices
	// a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// b := a[:]   // slice of all elements
	// c := a[3:]  // slice 4th elemnt to end
	// d := a[:6]  // slice first 6 elements
	// e := a[3:6] // slice 4th,5th and 6th elements
	// a[5] = 42
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	//? make in go
	// a := make([]int, 3, 100)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 2, 3, 4, 5, 6)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	//* play slice
	a := []int{1, 2, 3, 4, 5}
	// b := a[1:]
	// c := a[:len(a)-1]
	// fmt.Println(a, b, c)
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a) // creates random behaviour because of references
}
