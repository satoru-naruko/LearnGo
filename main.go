package main

import (
	"HelloWorld/mylib/my_database"
	"HelloWorld/mylib/say_hello"
	"fmt"
	"log"
)

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

	fmt.Println("** using package api **")
	say_hello.SayHello("taro")

	db, err := my_database.New("./mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted ID:", id)

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name string
		err := rows.Scan(&id, &name) // カラムの数と型に合わせて変数を用意
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID:", id, "Name:", name)
	}
}
