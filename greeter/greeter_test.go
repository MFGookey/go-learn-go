package main

import (
	"fmt"
	"testing"
)

func Test_Greet_BlankName(t *testing.T) {
	result, err := Greet("")
	if err == nil {
		t.Errorf("Blank name greeting was incorrect.  Expected error, but got \"%s\"", result)
	}
}

func Test_Greet_WithName(t *testing.T) {
	nameToGreet := "aslkjasdflkjh"
	expectedResult := fmt.Sprintf("Hello, %s", nameToGreet)
	result, err := Greet(nameToGreet)

	if result != expectedResult {
		t.Errorf("Greeting with name \"%s\" was incorrect.  Expected \"%s\" but got \"%s\"", nameToGreet, expectedResult, result)
	}

	if err != nil {
		t.Errorf("Greeting with name \"%s\" was incorrect.  Expected \"%s\" but got error: %s", nameToGreet, expectedResult, err)
	}
}

func Test_GreetWithGreeting_BlankName_BlankGreeting(t *testing.T) {
	greetingToUse := ""
	nameToUse := ""

	result, err := GreetWithGreeting(greetingToUse, nameToUse)

	if err == nil {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected error but got \"%s\"", result)
	}
}

func Test_GreetWithGreeting_Name_BlankGreeting(t *testing.T) {
	greetingToUse := ""
	nameToUse := "a;sdkjfafd;h"
	result, err := GreetWithGreeting(greetingToUse, nameToUse)
	if err == nil {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected error but got \"%s\"", result)
	}
}

func Test_GreetWithGreeting_BlankName_Greeting(t *testing.T) {
	greetingToUse := "laisdufhalwiuef"
	nameToUse := ""
	result, err := GreetWithGreeting(greetingToUse, nameToUse)
	if err == nil {
		t.Errorf("Greet with greeting with blank name and blank greeting was incorrect.  Expected error but got \"%s\"", result)
	}
}

func Test_GreetWithGreeting_Name_Greeting(t *testing.T) {
	greetingToUse := "asasfafs"
	nameToUse := "ASFafsaFS"
	expectedResult := fmt.Sprintf("%s, %s", greetingToUse, nameToUse)
	result, err := GreetWithGreeting(greetingToUse, nameToUse)

	if err != nil {
		t.Errorf("Greet with greeting with name and greeting was incorrect.  Expected \"%s\" but got error \"%s\"", expectedResult, err)
	}

	if result != expectedResult {
		t.Errorf("Greet with greeting with name and greeting was incorrect.  Expected \"%s\" but got \"%s\"", expectedResult, result)
	}
}
