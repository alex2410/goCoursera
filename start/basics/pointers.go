package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // новый указатель на переменную a

	fmt.Println(&a, a, b, *b, c, *c)

	// получение указателя на переменнут типа int
	// инициализировано значением по-умолчанию
	d := new(int)
	fmt.Println(d)
	*d = 12
	fmt.Println(d)
	*c = *d // c = 12 -> a = 12
	fmt.Println(&a, a, b, *b, c, *c, d, *d)
	*d = 13 // c и a не изменились
	fmt.Println(&a, a, b, *b, c, *c, d, *d)
	c = d   // теперь с указывает туда же, куда d
	*c = 14 // с = 14 -> d = 14, a = 12
	fmt.Println(&a, a, b, *b, c, *c, d, *d)

	*d = 16
	fmt.Println(&a, a, b, *b, c, *c, d, *d)
}
