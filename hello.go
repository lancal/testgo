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

/*func main() {

	var x [5]float64
	//x := [5]float64{ 98, 93, 77, 82, 83 }
	//z := make([]float64, 5)
	//var total float64 = 0

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

		for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)


		for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)) )
	fmt.Println("printed with len()")


		for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))


}*/

	//Slice

/* func main () {

	/*
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	*/

	/*
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

*/

func main () {

	//maps

	/*z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z["key"])
	delete(z, "key")
	fmt.Println(z["key"])*/

    /*

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])*/

	/*
	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}


	//name, ok := elements["O"]
	//fmt.Println(name, ok)

	if name, ok := elements["O"]; ok {
		fmt.Println(name, ok)
	}

	*/

	elements := map[string]map[string]string{

		"H": {
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": {
			"name":"Helium",
			"state":"gas",
		},
		"Li": {
			"name":"Lithium",
			"state":"solid",
		},

		"Be": {
			"name":"Beryllium",
			"state":"solid",
		},

		"B":  {
			"name":"Boron",
			"state":"solid",
		},
		"C":  {
			"name":"Carbon",
			"state":"solid",
		},
		"N":  {
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  {
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  {
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  {
			"name":"Neon",
			"state":"gas",
		},
	}
	if el, ok := elements["Li"]; ok {

		fmt.Println(el["name"], el["state"])
	}

}





