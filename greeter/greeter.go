package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	name := getUsersName()
	greeting, err := Greet(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}
}

func getUsersName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Whom shall we greet?")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	name = strings.Replace(name, "\r", "", -1)
	return name
}

//Greet Greet a person with Hello
func Greet(name string) (string, error) {
	return GreetWithGreeting("Hello", name)
}

//GreetWithGreeting Greet a person with a specific greeting
func GreetWithGreeting(greeting string, name string) (string, error) {
	if name != "" && greeting != "" {
		return fmt.Sprintf("%s, %s", greeting, name), nil
	} else {
		if name == "" {
			return "", errors.New("name cannot be blank")
		}

		return "", errors.New("greeting cannot be blank")
	}
}
