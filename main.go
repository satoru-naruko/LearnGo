package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var age int = 30
	fmt.Println("age:", age)

	var weight float64 = 70.5
	fmt.Println("weight:", weight)

	var name string = "Go lang"
	fmt.Println("name:", name)

	age_type_interface := 30
	fmt.Println("age_type_interface:", age_type_interface)

	var a = 10
	var b = 5

	fmt.Println("**comparison operator**")
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false

	fmt.Println("** conditional branch **")
	var day int = 3
	if day == 1 || day == 7 {
		fmt.Println("Holiday") // 休日
	} else if day >= 2 && day <= 6 {
		fmt.Println("Weekday") // 平日
	} else {
		fmt.Println("Invalid day") // 不正な曜日
	}
}
