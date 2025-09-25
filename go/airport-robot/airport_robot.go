package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greet Greeter) string {
	return greet.Greet(name)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	language := i.LanguageName()
	return fmt.Sprintf("I can speak %s: Ciao %s!", language, name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	language := p.LanguageName()
	return fmt.Sprintf("I can speak %s: Olá %s!", language, name)
}
