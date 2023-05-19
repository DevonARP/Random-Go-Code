package main

import "fmt"

func main() {

	score := 4
	var valid bool = true

	if score > 3 && !valid {
		fmt.Println("good")
	} else if score == 4 && valid {
		fmt.Println("good2")
	} else {
		fmt.Println("low")
	}
}
