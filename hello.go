package main

import "fmt"

/*func main() {
	fmt.Printf("hello, world\n")
}*/


/*func main(){

	fmt.Println("1 + 1 = ", 1 + 1)

}*/

/*func main() {
    fmt.Println(len("Hello World"))
    fmt.Println("Hello World"[1])
    fmt.Println("Hello " + "World")
}*/

/*func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}*/

/*
func main() {
	var x string = "Hello World"
	fmt.Println(x)
}*/

/*func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
}*/

/*unc	main(){

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}*/

/*func main ()  {

	var x string = "hello"
	var y string = "world"
	fmt.Println(x == y)

}*/

/*func main ()  {

	x := "Hello World"
	println(x)

}*/

/*func main ()  {

	var x = "Hello World"
	fmt.Println(x)

}*/

/*var ( a = 5 ; b = 10 ; c = 5 )

var(

	x1 = 12

	y = 14

	z = 15

)

*/

/*func main() {
	const x = "Hello World as constant"
	fmt.Println(x)

	fmt.Println(a , b , c)

	fmt.Println(x1, y , z)
}*/

/*func main() {

	var input float64

	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}*/

/*func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i=i+1 }
}*/

/*func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}*/

/*func main ()  {

	for i := 1; i <= 10 ;i++  {

		if i % 2 == 0{

			fmt.Println(i, "even")
		}else{

			fmt.Println(i, "odd")
		}

	}

}*/

// Array

/*func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])
}*/

func main() {

	var x [5]float64
	//x := [5]float64{ 98, 93, 77, 82, 83 }
	var total float64 = 0

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

/*	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
*/


/*	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)) )
	fmt.Println("printed with len()")
*/

	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))


}





