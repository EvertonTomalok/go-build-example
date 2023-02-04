package domain

import "fmt"

type Person struct {
	Name     string
	LastName string
}

func (p *Person) Greeting() {
	if p.Name == "" && p.LastName == "" {
		fmt.Println("You must provide your name via CLI. (--name and --lastName args)")
		return
	}
	fmt.Printf("Hello %s %s. How are you? \n", p.Name, p.LastName)
}
