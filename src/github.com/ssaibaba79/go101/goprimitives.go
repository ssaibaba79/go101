// Package definition
package main

//imports
import (
	"fmt"
	"strconv"
)

// global variables block
var (
	name   string = "Saravanan"
	email  string = "s@j.com"
	age    int    = 40
	senior bool   = false
)

// A global variable
var g int = 1000

// ExportVariable  variable visible outside of the package
var ExportVariable int = 1000

var u int

//function main
func main() {
	//Method scope variable
	var g int = 100

	fmt.Println(name, email, age, g)
	//fmt.Printf("%v, %T", u, u)
	printPrimitiveDefaults()
	arithmeticOps()
	booleanOps()
	bitOps()
	complexNumbers()
	stringOps()
	fmt.Println("g =", g, "string(g) =", string(g), "strconv.Itoa(g) =", strconv.Itoa(g))

}

func arithmeticOps() {

	//type is inferred
	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a + b)

	var c int8 = 120
	var d int16 = 32000

	// not allowed and c and d are not of same type int != int8
	//var x int8 = c + d
	// explicit cast is needed
	fmt.Println("Type casting : c(int8)=", c, "d(int16)=", d, "(c + int8(d))=", (c + int8(d)))

}

func booleanOps() {
	t := 1 == 1
	f := 1 == 2

	fmt.Printf("%T %v\n%T %v\n", t, t, f, f)
}

func bitOps() {
	a := 10 //1010
	b := 3  //0011

	fmt.Println("Bitwise AND : a=", a, "b=", b, "a & b =", a&b)    //0010
	fmt.Println("Bitwise OR  : a=", a, "b=", b, "a | b =", a|b)    //1011
	fmt.Println("Bitwise XOR : a=", a, "b=", b, "a ^ b =", a^b)    //1001
	fmt.Println("Bitwise NAND : a=", a, "b=", b, "a &^ b =", a&^b) //0100

	c := 8
	fmt.Println("Bit shift << : c=", c, "c << 3 =", c<<3) //2^3 * 2^3 = 2^6
	fmt.Println("Bit shift >> : c=", c, "c >> 3 =", c>>3) //2^3 / 2^3 = 2^0

}

// another function
func printPrimitiveDefaults() {

	//Type inferred initializtion
	// Can be use only in function scope, cannot be used in global scope
	FORMAT := "%T %v \n"

	// default value of all numeric types is 0
	// when size of the int is unspecified atleast 32 bit is
	// gauranteed based on the system
	var i int

	// signed integers
	// -128 to 127
	var i8 int8 = 126
	//  -32768 32767
	var i16 int16 = 32767
	// -2147483648 to 2147483647
	var i32 int32 = 2147483647
	// -9223372036854775808 to 9223372036854775807
	var i64 int64 = 9223372036854775807

	// byte is an alias of unint8
	var bite byte = 5
	//unsigned int
	// 0- 255
	var ui8 uint8 = 255
	// 0- 65535
	var ui16 uint16 = 65535
	// 0 - 4294967295
	var ui32 uint32 = 4294967295

	//unsigned 8 bit int

	var f float32
	var d float64

	// default value is false
	var b bool

	// default value is empty string
	var s string

	fmt.Printf(FORMAT, i, i)
	fmt.Printf(FORMAT, i8, i8)
	fmt.Printf(FORMAT, i16, i16)
	fmt.Printf(FORMAT, i32, i32)
	fmt.Printf(FORMAT, i64, i64)

	fmt.Printf(FORMAT, ui8, ui8)
	fmt.Printf(FORMAT, ui16, ui16)
	fmt.Printf(FORMAT, ui32, ui32)

	fmt.Printf(FORMAT, bite, bite)

	fmt.Printf(FORMAT, f, f)
	fmt.Printf(FORMAT, d, d)
	fmt.Printf(FORMAT, b, b)
	fmt.Printf(FORMAT, s, s)

}

func complexNumbers() {
	var n complex64 = 1 + 2i

	fmt.Printf("%T, %v\n", n, n)
	fmt.Printf("%T, %v\n", real(n), real(n))
	fmt.Printf("%T, %v\n", imag(n), imag(n))

	var m complex64 = 1 + 2i
	fmt.Printf("%T, %v\n", m, m)
	fmt.Printf("%T, %v\n", real(m), real(m))
	fmt.Printf("%T, %v\n", imag(m), imag(m))

	var cmplxNum = complex(5, 12)
	fmt.Printf("%T, %v\n", cmplxNum, cmplxNum)

}

func stringOps() {

	// utf8 chars
	var s string = "This is a string"

	// not allowed as string are immutable
	//s[2] = "u"
	fmt.Printf("%T, %v\n", s[5], s[5])
	fmt.Printf("%T, %v\n", string(s[5]), string(s[5]))

	// string concat
	var s1 string = " and another string"
	fmt.Printf("%T, %v\n", s+s1, s+s1)

	// string as bytes
	b := []byte(s)
	fmt.Printf("%T, %v\n", b, b)

	// rune is an alias for int32
	// utf32 character
	var r rune = 'a'
	fmt.Printf("%T, %v\n", r, r)

}
