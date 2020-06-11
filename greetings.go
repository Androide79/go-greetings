package greetings

import (
	"fmt"
)

// func main() {
// 	var emptyName string = ""
// 	var name string = "Pigi"

// 	fmt.Println(greetings(emptyName)) // this will return "Hello Dude!"
// 	fmt.Println(greetings(name)) // this will return "Hello Pigi!"
// }

func Greet(name string) string {
	var result string

	if len(name) == 0 {
		result = "Hello Dude!"
	} else {
		result = fmt.Sprintf("Hello %v!", name)
	}

	return result
}