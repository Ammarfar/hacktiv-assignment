package main

import (
	"assignment/helpers"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	user := helpers.GenerateUser(5)

	if input > len(user) {
		panic("User Not Found")
	}

	fmt.Printf("%+v\n", user[input-1])
}
