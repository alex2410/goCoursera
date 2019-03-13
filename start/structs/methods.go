package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры, для который вызван
func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetId(id int) {
	p.Id = id
}

// изменяет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl MySlice) Print() {
	fmt.Println(sl)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	pers1 := &Person{1, "Vasily"}
	pers := Person{1, "Vasily"}
	pers1.SetName("Vasily Romanov")
	(&pers).SetName("Vasily Romanov")
	pers.SetId(11)
	fmt.Printf("updated person: %#v\n", pers)
	fmt.Printf("updated person: %#v\n", pers1)

	var acc Account = Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Id:   2,
			Name: "Vasily Romanov",
		},
	}

	acc.SetName("romanov.vasily")
	acc.Person.SetName("Test")

	fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	sl.Print()
	fmt.Println(sl.Count(), sl)
}
