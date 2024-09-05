package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", g.LanguageName(), g.Greet(name))
}

type Italian struct {
}

func (i Italian) Greet(a string) string {
	return fmt.Sprintf("Ciao %s!", a)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct {
}

func (i Portuguese) Greet(a string) string {
	return fmt.Sprintf("Olá %s!", a)
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
