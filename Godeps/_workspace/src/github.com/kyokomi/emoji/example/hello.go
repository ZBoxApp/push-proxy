package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	fmt.Println("Hello Wolrd Emoji!")

	emoji.Println(":beer: Beer!!!")

	pizzaMessage := emoji.Sprint("I like a :pizza: and :sushi:!!")
	fmt.Println(pizzaMessage)
}
