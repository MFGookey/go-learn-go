package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	name := getUsersName()
	fmt.Println(Greet(name))
}

func getUsersName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Whom shall we greet?")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}

//Greet Greet a person with Hello
func Greet(name string) string {
	return GreetWithGreeting("Hello", name)
}

//GreetWithGreeting Greet a person with a specific greeting
func GreetWithGreeting(greeting string, name string) string {
	return fmt.Sprintf("%s, %s", greeting, name)
}
