package main

import (
	"fmt"
	"testing"
)

func Test_Greet_BlankName(t *testing.T) {
	blankGreeting := "Hello, "
	result := Greet("")
	if result != blankGreeting {
		t.Errorf("Blank name greeting was incorrect.  Expected \"%s\", but got \"%s\"", blankGreeting, result)
	}
}

func Test_Greet_WithName(t *testing.T) {
	nameToGreet := "aslkjasdflkjh"
	expectedResult := fmt.Sprintf("Hello, %s", nameToGreet)
	result := Greet(nameToGreet)

	if result != expectedResult {
		t.Errorf("Greeting with name \"%s\" was incorrect.  Expected \"%s\" but got \"%s\"", nameToGreet, expectedResult, result)
	}
}

func Test_GreetWithGreeting_BlankName_BlankGreeting(t *testing.T) {
	greetingToUse := ""
	nameToUse := ""
	expectedResult := fmt.Sprintf("%s, %s", greetingToUse, nameToUse)
	result := GreetWithGreeting(greetingToUse, nameToUse)
	if result != expectedResult {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected \"%s\" but got \"%s\"", expectedResult, result)
	}
}

func Test_GreetWithGreeting_Name_BlankGreeting(t *testing.T) {
	greetingToUse := ""
	nameToUse := "a;sdkjfafd;h"
	expectedResult := fmt.Sprintf("%s, %s", greetingToUse, nameToUse)
	result := GreetWithGreeting(greetingToUse, nameToUse)
	if result != expectedResult {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected \"%s\" but got \"%s\"", expectedResult, result)
	}
}

func Test_GreetWithGreeting_BlankName_Greeting(t *testing.T) {
	greetingToUse := "laisdufhalwiuef"
	nameToUse := ""
	expectedResult := fmt.Sprintf("%s, %s", greetingToUse, nameToUse)
	result := GreetWithGreeting(greetingToUse, nameToUse)
	if result != expectedResult {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected \"%s\" but got \"%s\"", expectedResult, result)
	}
}

func Test_GreetWithGreeting_Name_Greeting(t *testing.T) {
	greetingToUse := "asasfafs"
	nameToUse := "ASFafsaFS"
	expectedResult := fmt.Sprintf("%s, %s", greetingToUse, nameToUse)
	result := GreetWithGreeting(greetingToUse, nameToUse)
	if result != expectedResult {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected \"%s\" but got \"%s\"", expectedResult, result)
	}
}
