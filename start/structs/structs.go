package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	//embedded
	Person
}

type Accounts struct {
	accounts []*Account
}

func main() {
	// полное объявление структуры
	var acc Account = Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Name:    "Василий",
			Address: "Москва",
		},
	}

	var accounts Accounts = Accounts{}

	accounts.accounts = append(accounts.accounts, &acc)
	fmt.Printf("%#v\n", acc)
	fmt.Println()
	fmt.Println(accounts.accounts[0])
	acc.Name = "Name"
	fmt.Println(accounts.accounts[0])

	// короткое объявление структуры
	acc.Owner = Person{2, "Romanov Vasily", "Moscow"}

	fmt.Printf("%#v\n", acc)
	fmt.Println(accounts.accounts[0])

	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
}
